package fileman

import "fmt"

type FileManager interface {
	GetDirectories(path string) ([]string, error)
}

func NewFileMan(name string) (FileManager, error) {
	switch name {
	case "fs":
		return new(FileManFs), nil
	}
	return nil, fmt.Errorf("Unknown File Manager type: %s", name)
}
