package useCases

type UseCasesInteractors struct {
	fm  FileManager
	pom PlayOrderManager
	sp  SongPlayer
}

func NewUseCasesInteractors(fm FileManager, pom PlayOrderManager, sp SongPlayer) *UseCasesInteractors {
	uci := &UseCasesInteractors{fm, pom, sp}
	return uci
}
