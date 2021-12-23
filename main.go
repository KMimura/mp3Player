package main

import (
	"mp3Player/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	content := widgets.CreateMainWindow(myApp)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 450))
	myWindow.ShowAndRun()
}
