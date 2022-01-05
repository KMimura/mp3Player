package useCases

import (
	"mp3Player/entities"
)

type SongPlayer interface {
	Play(entities.Song)
	Stop(entities.Song)
	Next(entities.Song)
}
