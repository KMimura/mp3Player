package interfaceAdapters

import (
	"fmt"
	"io"
	"io/ioutil"
	"mp3Player/entities"
	"os"
	"strings"

	"github.com/dhowden/tag"
)

type FileManagerImpl struct{}

func (*FileManagerImpl) GetSongPathes(path string) []string {
	// get song pathes from a certain directory
	var filePathes []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return filePathes
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileNameSplitted := strings.Split(file.Name(), ".")
		if len(fileNameSplitted) <= 1 {
			continue
		}
		var acceptanceOnExtentionFlg bool
		for _, e := range acceptedExtentions {
			if fileNameSplitted[len(fileNameSplitted)-1] == e {
				acceptanceOnExtentionFlg = true
				break
			}
		}
		if acceptanceOnExtentionFlg {
			filePathes = append(filePathes, file.Name())
		}
	}
	return filePathes

}

func (*FileManagerImpl) GetASongFromPath(path string) *entities.Song {
	// create the song entity from a path
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
