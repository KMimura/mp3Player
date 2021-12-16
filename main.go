package main

import (
	// "image/color"

	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

// var grey = &color.Gray{Y: 123}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewButton("aaa", func() {
		fmt.Println("aaa")
	}))
	myWindow.Resize(fyne.NewSize(255, 510))
	// myWindow.SetContent(widget.NewLabel("Hello World!"))
	myWindow.ShowAndRun()
}
