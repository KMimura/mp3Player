package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateMainWindow(myApp fyne.App) *fyne.Container {
	appContainer := container.New(layout.NewVBoxLayout(), createMainScreen())
	return appContainer
}

func createMainScreen() *fyne.Container {
	songName := widget.NewLabel("SongName")
	songArtist := widget.NewLabel("SongArtist")
	songAlbum := widget.NewLabel("SongAlbum")
	songInfoContainer := container.New(layout.NewVBoxLayout(), songArtist, songAlbum)
	mainScreenContainer := container.New(layout.NewVBoxLayout(), songName, songInfoContainer)
	rectColorSheet := canvas.NewRectangle(color.White)
	rectContainer := container.New(layout.NewMaxLayout(), rectColorSheet, mainScreenContainer)
	return rectContainer
}
