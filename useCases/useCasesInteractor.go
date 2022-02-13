package useCases

type UseCasesInteractors struct {
	fm  FileManager
	pom PlayOrderManager
	sp  SongPlayer
}

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
