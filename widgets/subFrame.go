package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

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
