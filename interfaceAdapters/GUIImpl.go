package interfaceAdapters

import (
	"fmt"
	"image/color"
	"mp3Player/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type GUIImpl struct{}

// go the icons from https://icooon-mono.com/

func (*GUIImpl) ShowUserInterface() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	content := container.New(layout.NewVBoxLayout(), createMainScreen(), createButtons())
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 450))
	myWindow.ShowAndRun()
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
	leftImg := canvas.NewImageFromFile(utils.LeftArrowPic)
	leftImg.FillMode = canvas.ImageFillContain
	leftImg.SetMinSize(fyne.Size{Width: 50, Height: 50})
	leftButtonContainer := container.New(layout.NewMaxLayout(), leftButton, leftButtonSheet, leftImg)
	centerButton := widget.NewButton("", func() {
		fmt.Println("center")
	})
	centerButtonSheet := canvas.NewRectangle(color.Transparent)
	centerImg := canvas.NewImageFromFile(utils.StartButtonPic)
	centerImg.FillMode = canvas.ImageFillContain
	centerImg.SetMinSize(fyne.Size{Width: 50, Height: 50})
	centerButtonContainer := container.New(layout.NewMaxLayout(), centerButton, centerButtonSheet, centerImg)
	rightButton := widget.NewButton("", func() {
		fmt.Println("right")
	})
	rightButtonSheet := canvas.NewRectangle(color.Transparent)
	rightImg := canvas.NewImageFromFile(utils.RightArrowPic)
	rightImg.FillMode = canvas.ImageFillContain
	rightImg.SetMinSize(fyne.Size{Width: 50, Height: 50})
	rightButtonContainer := container.New(layout.NewMaxLayout(), rightButton, rightButtonSheet, rightImg)
	buttonsContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), leftButtonContainer, centerButtonContainer, rightButtonContainer, layout.NewSpacer())
	return buttonsContainer
}

func showAnother(appInstance fyne.App) {
	subWin := appInstance.NewWindow("Shown later")
	subMainLabel := widget.NewLabel("sub")
	subButton := widget.NewButton("close", func() {
		subWin.Close()
	})
	subContent := container.New(layout.NewVBoxLayout(), subMainLabel, subButton)
	subWin.SetContent(subContent)
	subWin.Resize(fyne.NewSize(200, 200))
	subWin.Show()
}

func NewGUIImpl() *GUIImpl {
	gui := &GUIImpl{}
	return gui
}
