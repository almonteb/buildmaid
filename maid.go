package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/almonteb/buildmaid/config"
	"github.com/almonteb/buildmaid/fileman"
	"github.com/spf13/viper"

	"github.com/almonteb/buildmaid/util"
	"path"
	"sync"
	"time"
)

var (
	version   string
	buildTime string
	runConfig config.Config
	pretend   bool
)

func init() {
	log.Printf("Welcome to Build Maid %s [%s]!", version, buildTime)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("BM")

	viper.BindEnv("DEBUG")
	if viper.GetBool("DEBUG") {
		log.SetLevel(log.DebugLevel)
	}

	viper.BindEnv("PRETEND")
	if viper.GetBool("PRETEND") {
		pretend = true
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file: %s", err)
	}
	if err := viper.Unmarshal(&runConfig); err != nil {
		log.Panicf("Unable to decode config into struct, %v", err)
	}
	log.Debugf("Config: %+v", runConfig)
}

func main() {
	wg := sync.WaitGroup{}
	workers := make(chan bool, runConfig.Global.MaxWorkers)
	defer close(workers)

	for {
		log.Debugf("Beginning run [Pretend mode: %t]", pretend)
		for project, config := range runConfig.Paths {
			fm, err := fileman.NewFileMan(config)
			if err != nil {
				log.Panicf("Unable to create file manager %s", config.FileMan)
			}
			for _, branch := range config.Branches {
				workers <- true
				wg.Add(1)
				go processProject(project, config, branch, fm, func() {
					<-workers
					wg.Done()
				})
			}
		}
		wg.Wait()
		if runConfig.Global.OneTime {
			log.Info("Run complete, exiting")
			break
		}
		log.Debugf("Run complete, sleeping for %d seconds", runConfig.Global.Interval)
		time.Sleep(time.Duration(runConfig.Global.Interval) * time.Second)
	}
}

func processProject(project string, projectCfg config.Project, branchCfg config.Branch, fm fileman.FileManager, onComplete func()) {
	defer onComplete()
	log.Printf("Processing project: %s, branch: %s, config: %+v", project, branchCfg.Name, branchCfg)
	root := path.Join(projectCfg.Root, branchCfg.Name)
	builds, err := fm.GetBuilds(root)
	if err != nil {
		log.Fatalf("Failed to get builds: %+v", err)
	}

	log.Debugf("Found builds: %+v", builds)
	toRemove := util.GetRemovalCandidates(util.FilterIgnores(builds, projectCfg.Ignores), branchCfg.MaxBuilds)
	log.Debugf("Removal candidates: %+v", toRemove)

	if pretend {
		return
	}

	for _, build := range toRemove {
		if err := fm.Delete(build); err != nil {
			log.Warnf("Unable to delete build: %s", build)
		}
	}
}
