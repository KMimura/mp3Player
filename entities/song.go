package entities

import (
	"fmt"
)

type Song struct {
	SongName     string
	SongLocation string
	Artist       string
	Album        Album
	Length       Artist
}

func (*Song) Play() {
	fmt.Println("play")
}

func (*Song) Stop() {
	fmt.Println("stop")
}
