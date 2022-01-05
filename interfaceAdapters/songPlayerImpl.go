package interfaceAdapters

import (
	"fmt"
	"mp3Player/entities"
)

type SongPlayerImpl struct {
	song entities.Song
}

func (*SongPlayerImpl) Play(entities.Song) {
	fmt.Println("here comes the song play impl")
}

func (*SongPlayerImpl) Stop(entities.Song) {
	fmt.Println("here comes the song stop impl")
}
