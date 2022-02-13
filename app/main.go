package main

import (
	"fmt"
	"mp3Player/factories"
)

func main() {
	// implementation of the use case
	factory := factories.NewMp3PlayerFactory()
	_, _, uiImpl, fmImpl, uci := factory.CreateIFAdapters()
	song := fmImpl.GetASongFromPath("music/sample album/sample_song.mp3")
	fmt.Println(song.SongName)
	uiImpl.ShowUserInterface(uci)
}
