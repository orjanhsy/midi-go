package main

import (
	"log"
	"slices"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
	"midi/backend"
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

func DeviceMenu(wg *sync.WaitGroup) *widget.List {
	deviceNames := getCurrentDeviceNames()
	deviceMenuState := state.DeviceMenuState{
		Devices:          binding.BindStringList(&deviceNames),
		ConnectedDevices: []string{},
	}

	onDeviceButtonClicked := func(deviceName string) {
		if slices.Contains(deviceMenuState.ConnectedDevices, deviceName) {
			log.Printf("Already listening to %s\n", deviceName)
			return
		}
		log.Printf("Now listening to device: %s\n", deviceName)
		go func() {
			defer wg.Done()
			deviceMenuState.ConnectedDevices = append(deviceMenuState.ConnectedDevices, deviceName)
			backend.ListenForMidiInput(deviceName)
		}()
	}

	list := widget.NewListWithData(deviceMenuState.Devices,
		func() fyne.CanvasObject {
			return widget.NewButton("template", nil)
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			buttonLabel, err := di.(binding.String).Get()
			if err != nil {
				log.Printf("Failed to convert bound string to string\n")
				return
			}
			co.(*widget.Button).OnTapped = func() {
				wg.Add(1)
				onDeviceButtonClicked(buttonLabel)

				if co.(*widget.Button).Icon == nil {
					iconPath := "assets/connected.png"
					icon, err := fyne.LoadResourceFromPath(iconPath)
					if err != nil {
						log.Printf("Failed to locate connected icon at %s", iconPath)
					}
					co.(*widget.Button).SetIcon(icon)
				} else {
					co.(*widget.Button).SetIcon(nil)
				}
			}
			co.(*widget.Button).SetText(buttonLabel)
		},
	)
	return list
}

func main() {
	var wg sync.WaitGroup
	appl := app.New()
	win := appl.NewWindow("Midi Listener")

	onDeviceMenuClicked := func() { win.SetContent(DeviceMenu(&wg)) }
	onListenerButtonClicked := func() { return }
	onSettingsButtonClicked := func() { return }

	deviceMenuButton := widget.NewButton("Enheter", onDeviceMenuClicked)
	listenerButton := widget.NewButton("Visualiser", onListenerButtonClicked)
	settingsButton := widget.NewButton("Innstillinger", onSettingsButtonClicked)

	bottomBar := container.NewGridWithColumns(3, deviceMenuButton, listenerButton, settingsButton)

	content := container.NewBorder(nil, bottomBar, nil, nil, nil)

	win.SetContent(content)
	win.ShowAndRun()
	appl.Quit()
	log.Println("Program closed")
	wg.Wait()
}
