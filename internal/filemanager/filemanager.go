package filemanager

import (
	"os"

	"github.com/rs/zerolog/log"
)

type FileManager struct {
}

func (m *FileManager) MustCreate(filePath string) *os.File {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to create file in path: %s", filePath)
	}
	return f
}

func New() *FileManager {
	return &FileManager{}
}
