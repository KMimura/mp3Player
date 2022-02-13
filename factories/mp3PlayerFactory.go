package factories

import (
	"mp3Player/interfaceAdapters"
	"mp3Player/useCases"
)

type Mp3PlayerFactory struct{}

func (*Mp3PlayerFactory) CreateIFAdapters() useCases.UserInteraction {
	sp := interfaceAdapters.NewSongPlayerImpl(nil)
	pom := interfaceAdapters.NewPlayOrderManagerImpl()
	fm := interfaceAdapters.NewFileManagerImpl()
	uci := useCases.NewUseCasesInteractors(fm, pom, sp)
	gui := interfaceAdapters.NewGUIImpl(uci)
	return gui
}

func NewMp3PlayerFactory() *Mp3PlayerFactory {
	mpf := Mp3PlayerFactory{}
	return &mpf
}
