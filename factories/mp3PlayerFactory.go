package factories

import (
	"mp3Player/interfaceAdapters"
	"mp3Player/useCases"
)

type Mp3PlayerFactory struct{}

func (*Mp3PlayerFactory) CreateIFAdapters(defaultSongFileLocation string) (useCases.SongPlayer, useCases.PlayOrderManager, useCases.UserInteraction, useCases.FileManager) {
	sp := interfaceAdapters.NewSongPlayerImpl(nil)
	pom := interfaceAdapters.NewPlayOrderManagerImpl(defaultSongFileLocation)
	gui := interfaceAdapters.NewGUIImpl()
	fm := interfaceAdapters.NewFileManagerImpl()
	return sp, pom, gui, fm
}

func NewMp3PlayerFactory() *Mp3PlayerFactory {
	mpf := Mp3PlayerFactory{}
	return &mpf
}
