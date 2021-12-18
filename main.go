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
	myWindow.ShowAndRun()
}

func showAnother(appInstance fyne.App, buttonCh chan string) {
	subWin := appInstance.NewWindow("Shown later")
	subMainLabel := widget.NewLabel("sub")
	subButton := widget.NewButton("close", func() {
		subWin.Close()
	})
	subContent := container.New(layout.NewVBoxLayout(), subMainLabel, subButton)
	subWin.SetContent(subContent)
	subWin.Resize(fyne.NewSize(200, 200))
	subWin.Show()
	buttonCh <- "done"
}
