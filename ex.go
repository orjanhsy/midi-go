package main

import (
	"fmt"
	"time"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

func xd() {
	defer midi.CloseDriver()

	// Print all available ports
	for _, p := range midi.GetInPorts() {
		fmt.Println("Available input port:", p.String())
	}

	// Print all available ports
	for _, p := range midi.GetOutPorts() {
		fmt.Println("Available input port:", p.String())
	}

	in, err := midi.FindInPort("MIDI Out")
	if err != nil {
		fmt.Println("can't find VMPK")
		return
	}

	stop, err := midi.ListenTo(in, func(msg midi.Message, timestampms int32) {
		var bt []byte
		var ch, key, vel uint8
		switch {
		case msg.GetSysEx(&bt):
			fmt.Printf("got sysex: % X\n", bt)
		case msg.GetNoteStart(&ch, &key, &vel):
			fmt.Printf("starting note %s on channel %v with velocity %v\n", midi.Note(key), ch, vel)
		case msg.GetNoteEnd(&ch, &key):
			fmt.Printf("ending note %s on channel %v\n", midi.Note(key), ch)
		default:
			// ignore
		}
	}, midi.UseSysEx())
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	defer stop()

	time.Sleep(time.Second * 5)

	select {}
}
