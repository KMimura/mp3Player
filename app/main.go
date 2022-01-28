package main

import (
	"fmt"
	"mp3Player/factories"
)

func main() {
	// implementation of the use case
	factory := factories.NewMp3PlayerFactory()
	_, _, uiImpl := factory.CreateIFAdapters("")
	ch := make(chan string)
	uiImpl.ShowUserInterface(ch)
	select {
	case <-ch:
		fmt.Println("received")
	default:
		fmt.Println("do nothing")
	}
}
