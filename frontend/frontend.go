package frontend

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Ui struct {
	app  fyne.App
	win  fyne.Window
	name string
}

func CreateView() *Ui {
	var ui Ui
	ui.name = "wtff"
	ui.app = app.New()
	ui.win = ui.app.NewWindow("MIDI Listener")

	hello := widget.NewLabel("Hello FYNE!")

	ui.win.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome!")
		}),
	))

	return &ui
}

func (ui *Ui) Render() {
	ui.win.ShowAndRun()
}

func (ui *Ui) PrintName() {
	fmt.Print(ui.name)
}
