package useCases

import "fmt"

type UseCasesInteractors struct {
	fm  FileManager
	pom PlayOrderManager
	sp  SongPlayer
}

// ideally this should be on some config file, but i am too lazy
var defaultFolder = "music/sample album"

func (uci *UseCasesInteractors) Play() {
	// get the current song from pom
	// get the song file from fm
	// play the song with sp
}

func (uci *UseCasesInteractors) Next() {
	// if currently playing, call the pause func
	// ask pom for the next song
	// get the song from fm
	// play the song with sp
}

func (uci *UseCasesInteractors) Back() {
	// if currently playing, call the pause func
	// ask pom for the previous song
	// get the song from fm
	// play the song with sp
	pathes := uci.fm.GetSongPathes(defaultFolder)
	for _, p := range pathes {
		fmt.Println(p)
	}
}

func (uci *UseCasesInteractors) Pause() {
	// pause the song with sp
}

func (uci *UseCasesInteractors) UpdateQueue() {
	// update pom
	// call play func
}

func NewUseCasesInteractors(fm FileManager, pom PlayOrderManager, sp SongPlayer) *UseCasesInteractors {
	uci := &UseCasesInteractors{fm, pom, sp}
	return uci
}
