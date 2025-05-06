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

	BottomBar      *fyne.Container
	DeviceMenu     *fyne.Container
	ListenerScreen *fyne.Container
	SettingsScreen *fyne.Container
}

func (ui *Ui) Init(dms state.DeviceMenuState, ls state.ListenerState) {
	ui.app = app.New()
	ui.win = ui.app.NewWindow("Midi-Lytter")

	onDeviceMenuClicked := func() {
		log.Println("User clicked deviceMenuScreen")
		ui.RenderDeviceMenu()
	}

	onListenerClicked := func() {
		log.Println("User clicked listenerScreen")
		ui.RenderListenerScreen()
	}

	onSettingsClicked := func() {
		log.Println("User clicked settingsScreen")
		ui.RenderSettingsScreen()
	}

	ui.BottomBar = CreateAppBar(
		onDeviceMenuClicked,
		onListenerClicked,
		onSettingsClicked,
	)

	pref := ui.app.Preferences()
	ui.DeviceMenu = CreateDeviceMenu(dms, ui.BottomBar)
	ui.ListenerScreen = CreateListenerScreen(
		ls, ui.BottomBar, pref,
	)
	ui.SettingsScreen = CreateSettingsScreen(pref, ui.BottomBar)
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

func (ui *Ui) RenderListenerScreen() {
	ui.win.SetContent(ui.ListenerScreen)
	ui.win.Content().Refresh()
}

func (ui *Ui) RenderSettingsScreen() {
	ui.win.SetContent(ui.SettingsScreen)
	ui.win.Content().Refresh()
}
