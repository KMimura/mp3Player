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
	leftButtonSheet := canvas.NewRectangle(color.Transparent)
	leftImg := canvas.NewImageFromFile("pics/left.png")
	leftImg.FillMode = canvas.ImageFillContain
	leftImg.SetMinSize(fyne.Size{Width: 30, Height: 30})
	leftButtonContainer := container.New(layout.NewMaxLayout(), leftButton, leftButtonSheet, leftImg)
	centerButton := widget.NewButton("", func() {
		fmt.Println("center")
	})
	centerButtonSheet := canvas.NewRectangle(color.Transparent)
	centerImg := canvas.NewImageFromFile("pics/stop.png")
	centerImg.FillMode = canvas.ImageFillContain
	centerImg.SetMinSize(fyne.Size{Width: 30, Height: 30})
	centerButtonContainer := container.New(layout.NewMaxLayout(), centerButton, centerButtonSheet, centerImg)
	rightButton := widget.NewButton("", func() {
		fmt.Println("right")
	})
	rightButtonSheet := canvas.NewRectangle(color.Transparent)
	rightImg := canvas.NewImageFromFile("pics/right.png")
	rightImg.FillMode = canvas.ImageFillContain
	rightImg.SetMinSize(fyne.Size{Width: 30, Height: 30})
	rightButtonContainer := container.New(layout.NewMaxLayout(), rightButton, rightButtonSheet, rightImg)
	buttonsContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), leftButtonContainer, centerButtonContainer, rightButtonContainer, layout.NewSpacer())
	return buttonsContainer
}
