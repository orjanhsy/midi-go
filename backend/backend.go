package backend

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

func ListenForMidiInput(color binding.String) {
	defer midi.CloseDriver()

	in, err := drivers.InByName("VMPK Output")
	if err != nil {
		fmt.Printf("Failed to locate diver")
		return
	}

	stop, err := midi.ListenTo(in, func(msg midi.Message, timestampms int32) {
		var bt []byte
		var ch, key, velo uint8
		switch {
		case msg.GetSysEx(&bt):
			fmt.Printf("Got sysex: %X\n", bt)
		case msg.GetNoteStart(&ch, &key, &velo):
			fmt.Printf("Read note %s", midi.Note(key).Name())

			// TODO
			col, err := color.Get()
			if err != nil {
				fmt.Println("Error when getting string from bound data")
				return
			}
			newCol := col + "XD"
			color.Set(newCol)

			fmt.Printf("-> setting color to: %s\n", newCol)
		case msg.GetNoteEnd(&ch, &key):
			// TODO
		default:
			// TODO
		}
	})
	if err != nil {
		fmt.Printf("Failed to listen to inPort")
		return
	}

	defer stop()
	select {}
}
