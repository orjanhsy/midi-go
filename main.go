package main

import (
	"log"

	"fyne.io/fyne/v2/data/binding"

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
	ui := frontend.Ui{}

	deviceNames := getCurrentDeviceNames()

	dms := state.DeviceMenuState{
		Devices:          binding.BindStringList(&deviceNames),
		ConnectedDevices: make(map[string]func()),
	}

	ui.Init(dms)
	ui.RenderDeviceMenu() // initial ui

	ui.Run()

	log.Println("Program closed")
}
