package useCases

import (
	"mp3Player/entities"
)

type FileManager interface {
	GetSongPathes(string) []string
	GetASongFromPath(string) *entities.Song
}
