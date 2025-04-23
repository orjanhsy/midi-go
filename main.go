package main

import (
	"fmt"
	"sync"

	"midi/backend"
	"midi/frontend"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	color := "#000000"

	go func() {
		defer wg.Done()
		backend.ListenForMidiInput(&color)
	}()
	fmt.Printf("Backend up\n")

	wg.Wait()
}
