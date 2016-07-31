package fileman

import (
	"fmt"
	"github.com/almonteb/buildmaid/config"
)

type FileManager interface {
	GetBuilds(path string) ([]string, error)
	Delete(path string) error
}

func NewFileMan(config config.Project) (FileManager, error) {
	switch config.FileMan {
	case "fs":
		return new(FileManFs), nil
	case "s3":
		return NewFileManS3(config.S3Config.Access, config.S3Config.Secret, config.S3Config.Bucket, config.S3Config.Host)
	}
	return nil, fmt.Errorf("Unknown File Manager type: %s", config.FileMan)
}
