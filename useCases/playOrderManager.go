package useCases

import "mp3Player/entities"

type PlayOrderManager interface {
	GetCurrentSong() *entities.Song
	UpdateIndex(int)
}
