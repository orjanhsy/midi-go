package backend

import (
	"fmt"
	"log"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"
)

func selectDriver(name string) (drivers.In, error) {
	in, err := midi.FindInPort(name)
	if err != nil {
		return nil, err
	} else {
		return in, nil
	}
}

// func handleNoteStart(msg midi.Message, ch uint8, key uint8, velo uint8) {
func handleNoteStart(key uint8, midiPortName string) {
	log.Printf("Read note %s from %s\n", midi.Note(key).Name(), midiPortName)
	return
}

func ListenForMidiInput(portName string) func() {
	in, err := midi.FindInPort(portName)
	if err != nil {
		log.Fatal("Couldnt find port")
	}

	stop, err := midi.ListenTo(in, func(msg midi.Message, timestampms int32) {
		var bt []byte
		var ch, key, velo uint8

		switch {
		case msg.GetSysEx(&bt):
			fmt.Printf("Got sysex: %X\n", bt)
			in.Close()
		case msg.GetNoteStart(&ch, &key, &velo):
			// handleNoteStart(msg, ch, key, velo)
			handleNoteStart(key, in.String())
		case msg.GetNoteEnd(&ch, &key):
			// TODO
		default:
			// TODO
		}
	}, midi.UseSysEx())
	if err != nil {
		log.Print("Failed to listen to inPort")
		return nil
	}

	return stop
}
