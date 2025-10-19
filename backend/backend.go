package backend

import (
	"fmt"
	"image/color"
	"log"
	"midi/clrconv"

	"gitlab.com/gomidi/midi/v2"
)

func GetCurrentDeviceNames() []string {
	devices := midi.GetInPorts()
	deviceNames := make([]string, len(devices))
	for i := 0; i < len(devices); i++ {
		deviceNames[i] = devices[i].String()
	}
	return deviceNames
}

var onNoteRecieved func(color.RGBA, string)

func SetNoteRecievedHandler(handler func(color.RGBA, string)) {
	if handler != nil {
		log.Println("Set Note Recieved handler")
		onNoteRecieved = handler
	} else {
		log.Println("Cannot set onNoteRecieved to nil")
	}
}

func handleNoteStart(key uint8, midiPortName string) {
	if onNoteRecieved == nil {
		log.Printf("Read note %s from %s. No action perfermed as no handler has been passed.\n", midi.Note(key).Name(), midiPortName)
	} else {
		log.Printf("Read note %s from %s.\n", midi.Note(key).Name(), midiPortName)
		col, err := clrconv.GetRGBAFromNote(midi.Note(key).Name())
		if err != nil {
			log.Println("Failed to convert color.")
		}

		name := midi.Note(key).Name()
		onNoteRecieved(col, name)
	}
}

func ListenForMidiInput(portName string) (func(), error) {
	in, err := midi.FindInPort(portName)
	if err != nil {
		log.Println("[LISTENER] Could not find port")
		return nil, err
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
		return nil, err
	}

	return stop, nil
}
