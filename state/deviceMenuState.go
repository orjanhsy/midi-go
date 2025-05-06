package state

import (
	"slices"
	"time"

	"fyne.io/fyne/v2/data/binding"

	"midi/backend"
)

type DeviceMenuState struct {
	Devices          binding.ExternalStringList
	ConnectedDevices map[string]func()
}

func (dms *DeviceMenuState) Init() {
	devices := backend.GetCurrentDeviceNames()
	dms.Devices = binding.BindStringList(&devices)
	go func() {
		for {
			curr_dev := backend.GetCurrentDeviceNames()
			dms.refreshDeviceList(curr_dev)
			dms.cleanupDisconnectedDevices(curr_dev)
			time.Sleep(time.Second * 2)
		}
	}()
	dms.ConnectedDevices = make(map[string]func())
}

func (dms *DeviceMenuState) ConnectDevice(deviceName string) error {
	stop, err := backend.ListenForMidiInput(deviceName)
	if err != nil {
		return err
	}

	dms.ConnectedDevices[deviceName] = stop
	return nil
}

func (dms *DeviceMenuState) refreshDeviceList(dev []string) {
	dms.Devices.Set(dev)
}

func (dms *DeviceMenuState) cleanupDisconnectedDevices(dev []string) {
	for name, stop := range dms.ConnectedDevices {
		if !slices.Contains(dev, name) {
			stop()
			delete(dms.ConnectedDevices, name)
		}
	}
}
