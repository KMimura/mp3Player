package main

import (
	"mp3Player/factories"
	"mp3Player/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	// implementation of the use case
	factory := factories.NewMp3PlayerFactory()
	songPlayer, fileManager := factory.CreateIFAdapters("")
	content := widgets.CreateMainWindow(myApp, &songPlayer, &fileManager)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 450))
	myWindow.ShowAndRun()
}
