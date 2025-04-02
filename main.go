package main

import (
	"errors"
	"fmt"
	"log"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

func main() {
	defer midi.CloseDriver()

	// color := "#000000" // black as default

	in, err := drivers.InByName("VMPK Output")
	if err != nil {
		log.Fatal("Failed to read inPorts")
	}

	stop, err := midi.ListenTo(in, func(msg midi.Message, timestampms int32) {
		var bt []byte
		var ch, key, vel uint8
		switch {
		case msg.GetSysEx(&bt):
			fmt.Printf("got sysex: % X\n", bt)
		case msg.GetNoteStart(&ch, &key, &vel):
			// fmt.Printf("Got note <%s> of value <%d>", midi.Note(key).Name(), midi.Note(key))
			fmt.Printf("Read note %s ", midi.Note(key).Name())
			newColor := noteToColor(midi.Note(key).Name())
			readColor, err := getReadableColorFromHexa(newColor)
			if err != nil {
				fmt.Print("Color not found in table")
			}
			fmt.Printf("-> color set to %s\n", readColor)

		case msg.GetNoteEnd(&ch, &key):
			// TODO
		default:
			// TODO
		}
	})
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	defer stop()
	select {}
}

func noteToColor(note string) string {
	switch note {
	case "A": // A
		return "#FF0000" // Red
	case "B": // B
		return "#FFA500" // Orange
	case "C": // C
		return "#FFFF00" // Yellow
	case "D": // D
		return "#00FF00" // Green
	case "E": // E
		return "#0000FF" // Blue
	case "F": // F
		return "#4B0082" // Indigo
	case "G": // G
		return "#EE82EE" // Violet
	default:
		return ""
	}
}

func getReadableColorFromHexa(hexString string) (string, error) {
	colors := map[string]string{
		"#FF0000": "Red",
		"#00FF00": "Green",
		"#0000FF": "Blue",
		"#FFFF00": "Yellow",
		"#00FFFF": "Cyan",
		"#FF00FF": "Magenta",
		"#FFFFFF": "White",
		"#000000": "Black",
		"#808080": "Gray",
		"#800000": "Maroon",
		"#808000": "Olive",
		"#008000": "Dark Green",
		"#800080": "Purple",
		"#FFC0CB": "Pink",
		"#A52A2A": "Brown",
		"#FFD700": "Gold",
		"#F0F8FF": "Alice Blue",
		"#FAEBD7": "Antique White",
		"#EE82EE": "Violet",
	}

	color, ok := colors[hexString]

	if ok {
		return color, nil
	}
	return "", errors.New("color not found")
}
