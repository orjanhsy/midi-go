package state

import "fyne.io/fyne/v2/data/binding"

type DeviceMenuState struct {
	Devices          binding.ExternalStringList
	ConnectedDevices map[string]func()
}
