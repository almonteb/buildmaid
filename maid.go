package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"

	"github.com/almonteb/buildmaid/config"
	"github.com/almonteb/buildmaid/fileman"
	"time"
)

var (
	version   string
	buildTime string
	runConfig config.Config
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
	defer func() {
		close(workers)
	}()

	for {
		log.Debugf("Beginning run")
		for project, config := range runConfig.Paths {
			fm, err := fileman.NewFileMan(config.FileMan)
			if err != nil {
				log.Panicf("Unable to create file manager %s", config.FileMan)
			}
			for _, branch := range config.Branches {
				workers <- true
				wg.Add(1)
				go processProject(project, branch, fm, &wg, workers)
			}
		}
		wg.Wait()
		log.Debugf("Run complete, sleeping for %d seconds", runConfig.Global.Interval)
		time.Sleep(time.Duration(runConfig.Global.Interval) * time.Second)
	}
}

func processProject(project string, branchCfg config.Branch, fm fileman.FileManager, wg *sync.WaitGroup, workers chan bool) {
	defer func() {
		<-workers
		wg.Done()
	}()
	log.Printf("Processing project: %s, branch: %s, config: %+v", project, branchCfg.Name, branchCfg)
	dirs, err := fm.GetDirectories(branchCfg.Name)
	if err != nil {
		log.Errorf("Failed to list directories: %+v", err)
	}
	log.Debugf("Found entries: %+v", dirs)
}
