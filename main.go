package main

import (
	// "image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// var grey = &color.Gray{Y: 123}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	hello := widget.NewLabel("Hello Fyne!")
	mainLabel := widget.NewLabel("Hello")
	button := widget.NewButton("aaa", func() {
		hello.SetText("aaa")
		ch := make(chan string)
		go showAnother(myApp, ch)
	})
	content := container.New(layout.NewVBoxLayout(), mainLabel, button)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(255, 510))
	// myWindow.SetContent(widget.NewLabel("Hello World!"))
	myWindow.ShowAndRun()
}

func showAnother(appInstance fyne.App, buttonCh chan string) {
	win := appInstance.NewWindow("Shown later")
	win.SetContent(widget.NewButton("close", func() {
		win.Close()
	}))
	win.Resize(fyne.NewSize(200, 200))
	win.Show()
	buttonCh <- "done"
}
