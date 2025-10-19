package main

import (
	"image/color"
	"log"
	"midi/frontend"
	"midi/state"

	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
)

func main() {
	ui := frontend.Ui{}
	dms := state.DeviceMenuState{}
	dms.Init()

	ls := state.ListenerState{}
	ls.Init(color.RGBA{0, 0, 0, 255})

	ui.Init(dms, ls)
	ui.RenderDeviceMenu() // initial ui

	ui.Run()

	log.Println("Program closed")
}
