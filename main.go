package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
	"midi/frontend"
	"midi/state"
)

func getCurrentDeviceNames() []string {
	devices := midi.GetInPorts()
	deviceNames := make([]string, len(devices))
	for i := 0; i < len(devices); i++ {
		deviceNames[i] = devices[i].String()
	}
	return deviceNames
}

func main() {
	appl := app.New()
	win := appl.NewWindow("Midi Listener")

	prevChan := make(chan fyne.CanvasObject, 1)
	onBackClicked := func() {
		log.Print("Back button clicked")
		prev := <-prevChan
		win.SetContent(prev)
		win.Content().Refresh()
	}
	backButton := widget.NewButton("Tilbake", onBackClicked)
	deviceNames := getCurrentDeviceNames()
	dms := state.DeviceMenuState{
		Devices:          binding.BindStringList(&deviceNames),
		ConnectedDevices: make(map[string]func()),
	}

	onDeviceMenuClicked := func() {
		log.Print("Device menu clicked\n")
		prev := win.Content()
		prevChan <- prev

		content := frontend.CreateDeviceMenu(dms, backButton)
		win.SetContent(content)
		win.Content().Refresh()
	}

	onListenerClicked := func() { return }
	onSettingsClicked := func() { return }

	deviceMenuButton := widget.NewButton("Enheter", onDeviceMenuClicked)
	listenerButton := widget.NewButton("Visualiser", onListenerClicked)
	settingsButton := widget.NewButton("Innstillinger", onSettingsClicked)

	bottomBar := container.NewGridWithColumns(3, deviceMenuButton, listenerButton, settingsButton)

	content := container.NewBorder(nil, bottomBar, nil, nil, nil)

	win.SetContent(content)

	win.ShowAndRun()
	appl.Quit()
	log.Println("Program closed")
}
