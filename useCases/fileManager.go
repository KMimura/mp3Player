package useCases

import (
	"mp3Player/entities"
)

type FileManager interface {
	GetSongsFromPath(string) []entities.Song
	GetASongFromPath(string) entities.Song
}
