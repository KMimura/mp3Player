package factories

import (
	"mp3Player/interfaceAdapters"
	"mp3Player/useCases"
)

type Mp3PlayerFactory struct{}

func (*Mp3PlayerFactory) CreateIFAdapters() (useCases.SongPlayer, useCases.PlayOrderManager, useCases.UserInteraction, useCases.FileManager, useCases.UseCasesInteractors) {
	sp := interfaceAdapters.NewSongPlayerImpl(nil)
	pom := interfaceAdapters.NewPlayOrderManagerImpl()
	gui := interfaceAdapters.NewGUIImpl()
	fm := interfaceAdapters.NewFileManagerImpl()
	uci := useCases.NewUseCasesInteractors(fm, pom, sp)
	return sp, pom, gui, fm, *uci
}

func NewMp3PlayerFactory() *Mp3PlayerFactory {
	mpf := Mp3PlayerFactory{}
	return &mpf
}
