package widgets

import (
	"fmt"
	"image/color"
	"mp3Player/useCases"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// go the icons from https://icooon-mono.com/

func CreateMainWindow(myApp fyne.App, songPlayer useCases.SongPlayer, fileManager useCases.FileManager) *fyne.Container {
	appContainer := container.New(layout.NewVBoxLayout(), createMainScreen(songPlayer), createButtons(songPlayer))
	fmt.Println(appContainer.Size())
	return appContainer
}

func createMainScreen(songPlayer useCases.SongPlayer) *fyne.Container {
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

func createButtons(soingPlayer useCases.SongPlayer) *fyne.Container {
	leftButton := widget.NewButton("", func() {
		fmt.Println("left")
	})
	leftButtonSheet := canvas.NewRectangle(color.Transparent)
	leftImg := canvas.NewImageFromFile("../pics/left.png")
	leftImg.FillMode = canvas.ImageFillContain
	leftImg.SetMinSize(fyne.Size{Width: 50, Height: 50})
	leftButtonContainer := container.New(layout.NewMaxLayout(), leftButton, leftButtonSheet, leftImg)
	centerButton := widget.NewButton("", func() {
		fmt.Println("center")
	})
	centerButtonSheet := canvas.NewRectangle(color.Transparent)
	centerImg := canvas.NewImageFromFile("../pics/start.png")
	centerImg.FillMode = canvas.ImageFillContain
	centerImg.SetMinSize(fyne.Size{Width: 50, Height: 50})
	centerButtonContainer := container.New(layout.NewMaxLayout(), centerButton, centerButtonSheet, centerImg)
	rightButton := widget.NewButton("", func() {
		fmt.Println("right")
	})
	rightButtonSheet := canvas.NewRectangle(color.Transparent)
	rightImg := canvas.NewImageFromFile("../pics/right.png")
	rightImg.FillMode = canvas.ImageFillContain
	rightImg.SetMinSize(fyne.Size{Width: 50, Height: 50})
	rightButtonContainer := container.New(layout.NewMaxLayout(), rightButton, rightButtonSheet, rightImg)
	buttonsContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), leftButtonContainer, centerButtonContainer, rightButtonContainer, layout.NewSpacer())
	return buttonsContainer
}
