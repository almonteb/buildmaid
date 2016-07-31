package fileman

import (
	"io/ioutil"
	"os"
	"path"
)

type FileManFs struct{}

func (fm *FileManFs) GetBuilds(root string) ([]Build, error) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}

	var builds []Build
	for _, dir := range files {
		if dir.IsDir() {
			p := path.Join(root, dir.Name())
			builds = append(builds, Build{Name: dir.Name(), Path: p})
		}
	}
	return builds, nil
}

func (fm *FileManFs) Delete(build Build) error {
	return os.RemoveAll(build.Path)
}
