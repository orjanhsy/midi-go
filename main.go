package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
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
		ConnectedDevices: make(map[string]func()),
	}

	onDeviceClicked := func(deviceName string, co *widget.Button) {
		if stop, exists := deviceMenuState.ConnectedDevices[deviceName]; exists {
			stop()
			delete(deviceMenuState.ConnectedDevices, deviceName)
			co.Icon = nil
			log.Printf("Stopped listening to device %s\n", deviceName)
			return
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			stop := backend.ListenForMidiInput(deviceName)
			deviceMenuState.ConnectedDevices[deviceName] = stop
		}()

		log.Printf("Now listening to device: %s\n", deviceName)

		iconPath := "assets/connected.png"
		icon, err := fyne.LoadResourceFromPath(iconPath)
		if err != nil {
			log.Printf("Failed to locate connected icon at %s", iconPath)
		}
		co.SetIcon(icon)
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
				onDeviceClicked(buttonLabel, co.(*widget.Button))
			}
			co.(*widget.Button).SetText(buttonLabel)
		},
	)
	return list
}

func main() {
	var wg sync.WaitGroup
	defer midi.CloseDriver()
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

	onDeviceMenuClicked := func() {
		log.Print("Device menu clicked\n")
		prev := win.Content()
		prevChan <- prev

		content := container.NewBorder(nil, backButton, nil, nil, DeviceMenu(&wg))
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
	wg.Wait()
}
