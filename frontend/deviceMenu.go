package frontend

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"gitlab.com/gomidi/midi/v2"
	"midi/backend"
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

		stop := backend.ListenForMidiInput(deviceName)
		dms.ConnectedDevices[deviceName] = stop

		log.Printf("Now listening to device: %s\n", deviceName)

		iconPath := "assets/connected.png"
		icon, err := fyne.LoadResourceFromPath(iconPath)
		if err != nil {
			log.Printf("Failed to locate connected icon at %s", iconPath)
		}
		co.SetIcon(icon)
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

				iconPath := "assets/connected.png"
				icon, err := fyne.LoadResourceFromPath(iconPath)
				if err != nil {
					log.Printf("Failed to locate connected icon at %s", iconPath)
				}
				co.(*widget.Button).SetIcon(icon)
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
