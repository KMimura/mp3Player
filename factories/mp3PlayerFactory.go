package factories

import (
	"mp3Player/interfaceAdapters"
	"mp3Player/useCases"
)

type Mp3PlayerFactory struct{}

func (*Mp3PlayerFactory) CreateIFAdapters(defaultSongFileLocation string) (useCases.SongPlayer, useCases.PlayOrderManager, useCases.UserInteraction) {
	sp := interfaceAdapters.NewSongPlayerImpl(nil)
	pom := interfaceAdapters.NewFileManagerImpl(defaultSongFileLocation)
	gui := interfaceAdapters.NewGUIImpl()
	return sp, pom, gui
}

func NewMp3PlayerFactory() *Mp3PlayerFactory {
	mpf := Mp3PlayerFactory{}
	return &mpf
}
