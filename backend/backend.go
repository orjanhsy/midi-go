package backend

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/data/binding"

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

// func handleNoteStart(msg midi.Message, ch uint8, key uint8, velo uint8) {
func handleNoteStart(key uint8) {
	log.Printf("Read note %s\n", midi.Note(key).Name())
	return
}

func ListenForMidiInput(clrLab binding.String) {
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
			// handleNoteStart(msg, ch, key, velo)
			handleNoteStart(key)
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
