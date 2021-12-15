package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.Resize(fyne.NewSize(255, 510))
	myWindow.SetContent(widget.NewLabel("Hello World!"))
	myWindow.ShowAndRun()
}
