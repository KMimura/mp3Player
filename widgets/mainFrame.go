package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateMainWindow(myApp fyne.App) *fyne.Container {
	hello := widget.NewLabel("Hello Fyne!")
	mainLabel := widget.NewLabel("Hello")
	button := widget.NewButton("aaa", func() {
		hello.SetText("aaa")
		ch := make(chan string)
		go showAnother(myApp, ch)
	})
	content := container.New(layout.NewVBoxLayout(), mainLabel, button)
	return content
}
