package config

type global struct {
	MaxWorkers int  `mapstructure:"max_workers"`
	OneTime    bool `mapstructure:"one_time"`
	Interval   int
}

// Branch configuration
type Branch struct {
	Name    string
	MaxDays int `mapstructure:"max_days"`
}

// Project configuration
type Project struct {
	Root     string
	FileMan  string `mapstructure:"file_man"`
	Branches []Branch
}

// Config for config
type Config struct {
	Global global
	Paths  map[string]Project
}
