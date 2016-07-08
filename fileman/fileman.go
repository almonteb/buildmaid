package fileman

import "fmt"

type FileManager interface {
	GetDirectories(path string) ([]string, error)
	Delete(path string) error
}

type File struct {
	Name        string
	IsDirectory bool
}

func NewFileMan(name string) (FileManager, error) {
	switch name {
	case "fs":
		return new(FileManFs), nil
	}
	return nil, fmt.Errorf("Unknown File Manager type: %s", name)
}
