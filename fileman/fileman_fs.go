package fileman

import (
	"io/ioutil"
	"os"
	"path"
)

type FileManFs struct{}

func (fm *FileManFs) GetBuilds(root string) ([]string, error) {

	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, dir := range files {
		if dir.IsDir() {
			p := path.Join(root, dir.Name())
			dirs = append(dirs, p)
		}
	}
	return dirs, nil
}

func (fm *FileManFs) Delete(path string) error {
	return os.RemoveAll(path)
}
