package useCases

type PlayOrderManager interface {
	GetCurrentSong() string
	UpdateIndex(int)
}
