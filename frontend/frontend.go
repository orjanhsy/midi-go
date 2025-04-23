package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type Ui struct {
	app        fyne.App
	win        fyne.Window
	colorLabel *widget.Label
	clock      *widget.Label
}

func CreateView(color *binding.String) *Ui {
	var ui Ui
	ui.app = app.New()
	ui.win = ui.app.NewWindow("MIDI Listener")
	ui.colorLabel = widget.NewLabelWithData(*color)

	ui.win.SetContent(container.NewVBox(
		ui.colorLabel,
	))

	return &ui
}

func (ui *Ui) Start() {
	ui.win.ShowAndRun()
	ui.app.Quit()
}
