package main

import (
	"log"

	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
	"midi/frontend"
	"midi/state"
)

func main() {
	ui := frontend.Ui{}
	dms := state.DeviceMenuState{}
	dms.Init()

	ls := state.ListenerState{}
	ls.Init("black")

	ui.Init(dms, ls)
	ui.RenderDeviceMenu() // initial ui

	ui.Run()

	log.Println("Program closed")
}
