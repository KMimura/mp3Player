package interfaceAdapters

import (
	"io/ioutil"
	"strings"
)

var acceptedExtentions = []string{"mp3"}

type PlayOrderManagerImpl struct {
	queue []string
	index int
}

func (pom *PlayOrderManagerImpl) GetSongPathes(path string) []string {
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
	pom.index = 0
	return filePathes
}

func (pom *PlayOrderManagerImpl) UpdateIndex(index int) {
	pom.index++
	return
}

func (*PlayOrderManagerImpl) GetCurrentSong() string {
	return ""
}

func NewPlayOrderManagerImpl(path string) *PlayOrderManagerImpl {
	pom := &PlayOrderManagerImpl{queue: nil}
	var songs []string
	if path != "" {
		songs = pom.GetSongPathes(path)
	}
	pom.queue = songs
	return pom
}
