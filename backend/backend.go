package backend

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
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

var onNoteRecieved func(string)

func SetNoteRecievedHandler(handler func(string)) {
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
		name := midi.Note(key).Name()
		onNoteRecieved(name)
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

func SaveNoteColors(prefs fyne.Preferences, colors map[string]color.RGBA) {
	data, err := json.Marshal(colors)
	if err != nil {
		log.Println("Failed to marshal color map", err)
		return
	}
	prefs.SetString("noteColors", string(data))
}

func LoadNoteColors(prefs fyne.Preferences) map[string]color.RGBA {
	def := map[string]color.RGBA{
		"A":  {R: 255, G: 255, B: 0, A: 255},  // yellow
		"B":  {R: 64, G: 224, B: 208, A: 255}, // turquoise
		"C":  {R: 0, G: 0, B: 0, A: 255},      // black
		"D":  {R: 0, G: 0, B: 139, A: 255},    // really blue (dark blue)
		"Db": {R: 0, G: 70, B: 200, A: 255},
		"E":  {R: 0, G: 100, B: 0, A: 255},     // dark green
		"F":  {R: 255, G: 165, B: 0, A: 255},   // orange
		"Gb": {R: 144, G: 238, B: 144, A: 255}, // light green (renamed from duplicate "f")
		"G":  {R: 255, G: 0, B: 0, A: 255},     // red
	}

	data := prefs.String("noteColors")
	if data == "" {
		return def
	}
	var colors map[string]color.RGBA

	if err := json.Unmarshal([]byte(data), &colors); err != nil {
		log.Println("Failed to unmarshal color map", err)
		return def
	}

	return colors
}
