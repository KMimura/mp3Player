package factories

import (
	"mp3Player/interfaceAdapters"
	"mp3Player/useCases"
)

type Mp3PlayerFactory struct{}

func (*Mp3PlayerFactory) CreateIFAdapters(defaultSongFileLocation string) (useCases.SongPlayer, useCases.PlayOrderManager) {
	sp := interfaceAdapters.NewSongPlayerImpl(nil)
	pom := interfaceAdapters.NewFileManagerImpl(defaultSongFileLocation)
	return sp, pom
}

func NewMp3PlayerFactory() *Mp3PlayerFactory {
	mpf := Mp3PlayerFactory{}
	return &mpf
}
