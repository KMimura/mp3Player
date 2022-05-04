package interfaceAdapters

import (
	"mp3Player/entities"
)

var acceptedExtentions = []string{"mp3"}

type PlayOrderManagerImpl struct {
	queue []entities.Song
	index int
}

func (pom *PlayOrderManagerImpl) UpdateIndex(index int) {
	pom.index++
	return
}

func (*PlayOrderManagerImpl) GetCurrentSong() *entities.Song {
	return &entities.Song{}
}

func NewPlayOrderManagerImpl() *PlayOrderManagerImpl {
	pom := &PlayOrderManagerImpl{queue: nil}
	return pom
}
