package interfaceAdapters

import (
	"mp3Player/entities"
)

type FileManagerImpl struct {
	queue []*entities.Song
}

func (*FileManagerImpl) GetSongsFromPath(path string) []entities.Song {
	return nil
}

func (*FileManagerImpl) GetASongFromPath(path string) entities.Song {
	return entities.Song{}
}

func NewFileManagerImpl(q []*entities.Song) *FileManagerImpl {
	return &FileManagerImpl{queue: q}
}
