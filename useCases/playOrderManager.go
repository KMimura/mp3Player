package useCases

type PlayOrderManager interface {
	GetSongPathes(string) []string
	GetCurrentSong() string
	UpdateIndex(int)
}
