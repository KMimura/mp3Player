package main

import (
	"mp3Player/interfaceAdapters"
	"mp3Player/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	// implementation of the use case
	songPlayer := interfaceAdapters.SongPlayerImpl{}
	content := widgets.CreateMainWindow(myApp, &songPlayer)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 450))
	myWindow.ShowAndRun()
}