package state

import (
	"fyne.io/fyne/v2/data/binding"

	"midi/backend"
)

type DeviceMenuState struct {
	Devices          binding.ExternalStringList
	ConnectedDevices map[string]func()
}

func (dms *DeviceMenuState) ConnectDevice(deviceName string) {
	stop := backend.ListenForMidiInput(deviceName)
	dms.ConnectedDevices[deviceName] = stop
}
