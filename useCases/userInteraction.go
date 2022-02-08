package useCases

import (
	"mp3Player/entities"
)

type UserInteraction interface {
	ShowUserInterface()
	UpdatePlayingSong(*entities.Song)
}
