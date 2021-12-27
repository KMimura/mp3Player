package widgets

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateMainWindow(myApp fyne.App) *fyne.Container {
	appContainer := container.New(layout.NewVBoxLayout(), createMainScreen(), createButtons())
	fmt.Println(appContainer.Size())
	return appContainer
}

func createMainScreen() *fyne.Container {
	songName := canvas.NewText("Song Name", color.Black)
	songName.Alignment = 1
	songArtist := canvas.NewText("Song Artist", color.Black)
	songArtist.Alignment = 0
	songAlbum := canvas.NewText("Song Album", color.Black)
	songAlbum.Alignment = 0
	songInfoContainer := container.New(layout.NewVBoxLayout(), songArtist, songAlbum)
	mainScreenContainer := container.New(layout.NewVBoxLayout(), songName, songInfoContainer)
	rectColorSheet := canvas.NewRectangle(color.White)
	rectContainer := container.New(layout.NewMaxLayout(), rectColorSheet, mainScreenContainer)
	rectContainer.Resize(fyne.Size{Width: 200, Height: 200})
	return rectContainer
}

func createButtons() *fyne.Container {
	leftButton := widget.NewButton("", func() {
		fmt.Println("left")
	})
	centerButton := widget.NewButton("", func() {
		fmt.Println("center")
	})
	rightButton := widget.NewButton("", func() {
		fmt.Println("right")
	})
	buttonsContainer := container.New(layout.NewHBoxLayout(), leftButton, centerButton, rightButton)
	return buttonsContainer
}
