package backend

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"
	"midi/clrconv"
)

func selectDriver(name string) (drivers.In, error) {
	in, err := midi.FindInPort(name)
	if err != nil {
		return nil, err
	} else {
		return in, nil
	}
}

func readIndexFromStdin() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Skriv enhetens id: ")
	strIndex, err := reader.ReadString('\n')
	if err != nil {
		println("Failed to read keyboard input")
		return -1
	}

	index, err := strconv.Atoi(strings.TrimSuffix(strIndex, "\n"))
	if err != nil {
		fmt.Printf("Failed to convert <%s> to int. Got %d\n", strIndex, index)
		return -1
	}
	return index
}

func ListenForMidiInput(clrLab binding.String, colorChangeHandler func(string)) {
	defer midi.CloseDriver()
	inPorts := midi.GetInPorts()
	fmt.Print(inPorts.String())

	index := readIndexFromStdin()
	if index < 0 {
		fmt.Println("Reading index failed")
		return
	}

	in := inPorts[index]
	if in == nil {
		fmt.Println("Selected device was nil")
		return
	} else {
		fmt.Printf("Connected to %s\n", in.String())
	}

	stop, err := midi.ListenTo(in, func(msg midi.Message, timestampms int32) {
		var bt []byte
		var ch, key, velo uint8
		switch {
		case msg.GetSysEx(&bt):
			fmt.Printf("Got sysex: %X\n", bt)
		case msg.GetNoteStart(&ch, &key, &velo):
			fmt.Printf("Read note %s", midi.Note(key).Name())

			newColRGB := clrconv.NoteToRGBColor(midi.Note(key).Name())
			newCol, err := clrconv.GetReadableColorFromRGB(newColRGB)
			if err != nil {
				fmt.Print("Failed to convert color from rbg to readable.\n")
				return
			}
			fmt.Printf("-> setting color to: %s\n", newCol)

			fyne.Do(func() {
				colorChangeHandler(newCol)
			})

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
