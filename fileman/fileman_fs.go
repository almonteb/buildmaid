package fileman

type FileManFs struct{}

func (fm *FileManFs) GetDirectories(path string) ([]string, error) {
	return []string{"1", "2"}, nil
}
