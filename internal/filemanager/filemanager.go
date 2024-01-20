package filemanager

import "os"

type FileManager struct {
}

func (m *FileManager) MustCreate(filePath string) *os.File {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	return f
}

func New() *FileManager {
	return &FileManager{}
}
