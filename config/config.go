package config

type global struct {
	MaxWorkers int  `mapstructure:"max_workers"`
	OneTime    bool `mapstructure:"one_time"`
	Interval   int
}

// Branch configuration
type Branch struct {
	Name      string
	MaxBuilds int `mapstructure:"max_builds"`
}

// S3Config information
type S3Config struct {
	Access string
	Secret string
	Bucket string
	Host   string
}

// Project configuration
type Project struct {
	Root     string
	FileMan  string `mapstructure:"file_man"`
	Ignores  []string
	Branches []Branch
	S3Config S3Config `mapstructure:"s3_config"`
}

// Config for config
type Config struct {
	Global global
	Paths  map[string]Project
}
