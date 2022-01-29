package interfaceAdapters

import (
	"io"
	"mp3Player/entities"
	"os"

	"github.com/dhowden/tag"
)

type FileManagerImpl struct{}

func (*FileManagerImpl) GetASongFromPath(path string) *entities.Song {
	file, err := os.Open(path)
	readseek := io.ReadSeeker(file)
	if err != nil {
		// TODO handle the exception in a proper way
		panic(err)

	}
	songMetadata, err := tag.ReadFrom(readseek)
	if err != nil {
		// TODO handle the exception in a proper way
		panic(err)
	}
	songName := songMetadata.Title()
	songAlbum := songMetadata.Album()
	songArtist := songMetadata.Artist()
	songData := entities.Song{SongName: songName, SongLocation: path, Artist: songArtist, Album: songAlbum}
	return &songData
}
