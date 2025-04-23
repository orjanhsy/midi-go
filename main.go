package main

import (
	"fmt"
	"sync"

	"fyne.io/fyne/v2/data/binding"

	"midi/backend"
	"midi/frontend"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	color := binding.NewString()
	ui := frontend.CreateView(&color)
	color.Set("Yellow")

	go func() {
		defer wg.Done()
		backend.ListenForMidiInput(color)
	}()
	fmt.Printf("Backend up\n")

	ui.Start()

	wg.Wait()
}
