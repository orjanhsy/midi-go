package frontend

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"midi/state"
)

type Ui struct {
	app fyne.App
	win fyne.Window

	BottomBar  *fyne.Container
	DeviceMenu *fyne.Container
}

func (ui *Ui) Init(dms state.DeviceMenuState) {
	ui.app = app.New()
	ui.win = ui.app.NewWindow("Midi-Lytter")

	onDeviceMenuClicked := func() {
		log.Print("Device menu clicked\n")
		ui.RenderDeviceMenu()
	}

	onListenerClicked := func() { return }

	onSettingsClicked := func() { return }

	ui.BottomBar = CreateAppBar(
		onDeviceMenuClicked,
		onListenerClicked,
		onSettingsClicked,
	)

	ui.DeviceMenu = CreateDeviceMenu(dms, ui.BottomBar)
}

// blocks
func (ui *Ui) Run() {
	ui.win.ShowAndRun()
	ui.app.Quit()
}

func (ui *Ui) RenderDeviceMenu() {
	ui.win.SetContent(ui.DeviceMenu)
	ui.win.Content().Refresh()
}
