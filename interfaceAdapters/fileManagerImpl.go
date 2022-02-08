package interfaceAdapters

import (
	"fmt"
	"io"
	"mp3Player/entities"
	"os"
	"strings"

	"github.com/dhowden/tag"
)

type FileManagerImpl struct{}

func (*FileManagerImpl) GetASongFromPath(path string) *entities.Song {
	file, err := os.Open(path)
	readseek := io.ReadSeeker(file)
	if err != nil {
		// TODO handle the exception in a proper way
		fmt.Println("a")
		panic(err)

	}
	songMetadata, err := tag.ReadFrom(readseek)
	if err != nil {
		// TODO handle the exception in a proper way
		fmt.Println("b")
		panic(err)
	}
	songName := songMetadata.Title()
	if songName == "" {
		splitSongName := strings.Split(path, "/")
		lastElement := splitSongName[len(splitSongName)-1]
		lastElementSplit := strings.Split(lastElement, ".")
		var songNameConcat string
		for i := 0; i < len(lastElementSplit)-1; i++ {
			songNameConcat += lastElementSplit[i]
		}
		songName = songNameConcat
	}
	songAlbum := songMetadata.Album()
	songArtist := songMetadata.Artist()
	songData := entities.Song{SongName: songName, SongLocation: path, Artist: songArtist, Album: songAlbum}
	return &songData
}

func NewFileManagerImpl() *FileManagerImpl {
	fm := FileManagerImpl{}
	return &fm
}
