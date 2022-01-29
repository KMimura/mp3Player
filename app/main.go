package main

import (
	"mp3Player/factories"
)

func main() {
	// implementation of the use case
	factory := factories.NewMp3PlayerFactory()
	_, _, uiImpl := factory.CreateIFAdapters("")
	uiImpl.ShowUserInterface()

}
