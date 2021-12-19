package main

import (
	// "image/color"

	"mp3Player/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// var grey = &color.Gray{Y: 123}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	content := widgets.CreateMainWindow(myApp)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(255, 510))
	myWindow.ShowAndRun()
}
