package fileman

import (
	"io/ioutil"
	"os"
)

type FileManFs struct{}

func (fm *FileManFs) GetDirectories(path string) ([]string, error) {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, dir := range files {
		if dir.IsDir() {
			dirs = append(dirs, dir.Name())
		}
	}
	return dirs, nil
}

func (fm *FileManFs) Delete(path string) error {
	return os.RemoveAll(path)
}
