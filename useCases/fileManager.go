package useCases

import (
	"mp3Player/entities"
)

type FileManager interface {
	GetASongFromPath(string) *entities.Song
}
