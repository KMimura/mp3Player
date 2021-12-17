package main

import (
	// "image/color"

	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// var grey = &color.Gray{Y: 123}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	hello := widget.NewLabel("Hello Fyne!")
	myWindow.SetContent(container.NewVBox(
		hello,
		widget.NewButton("aaa", func() {
			hello.SetText("aaa")
			ch := make(chan string)
			go buttonFunc(ch)
		})))
	myWindow.Resize(fyne.NewSize(255, 510))
	// myWindow.SetContent(widget.NewLabel("Hello World!"))
	myWindow.ShowAndRun()
}

func buttonFunc(buttonCh chan string) {
	fmt.Println("aaa")
	buttonCh <- "done"
}
