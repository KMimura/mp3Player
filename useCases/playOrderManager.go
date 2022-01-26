package useCases

type PlayOrderManager interface {
	GetSongPathes(string) []string
	UpdateIndex(int)
}
