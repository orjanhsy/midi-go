package frontend

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"gitlab.com/gomidi/midi/v2"
	"midi/state"
)

func createDeviceList(dms state.DeviceMenuState) *widget.List {
	onDeviceClicked := func(deviceName string, co *widget.Button) {
		if stop, exists := dms.ConnectedDevices[deviceName]; exists {
			stop()
			midi.CloseDriver()
			delete(dms.ConnectedDevices, deviceName)
			co.SetIcon(nil)
			log.Printf("Stopped listening to device %s\n", deviceName)
			return
		}

		if err := dms.ConnectDevice(deviceName); err != nil {
			log.Printf("[DEVICE MENU] Could not connect to device")
			co.SetIcon(nil)
			return
		}
		log.Printf("Now listening to device: %s\n", deviceName)

		co.SetIcon(ResourceConnectedPng)
	}

	list := widget.NewListWithData(dms.Devices,
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

			if _, exists := dms.ConnectedDevices[buttonLabel]; exists {
				co.(*widget.Button).SetIcon(ResourceConnectedPng)
			} else {
				co.(*widget.Button).SetIcon(nil)
			}

			co.(*widget.Button).SetText(buttonLabel)
		},
	)
	return list
}

func CreateDeviceMenu(dms state.DeviceMenuState, bottomBar *fyne.Container) *fyne.Container {
	list := createDeviceList(dms)
	menu := container.NewBorder(nil, bottomBar, nil, nil, list)

	return menu
}
