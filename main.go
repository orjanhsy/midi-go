package main

import (
	"image/color"
	"sync"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
	"midi/backend"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	boundClrLab := binding.NewString()
	appl := app.New()
	win := appl.NewWindow("Midi Listener")

	boundClrLab.Set("White")
	clrLab := widget.NewLabelWithData(boundClrLab)
	rect := canvas.NewRectangle(color.White)

	go func() {
		defer wg.Done()
		backend.ListenForMidiInput(boundClrLab)
	}()

	win.SetContent(
		container.NewBorder(clrLab, nil, nil, nil, rect),
	)

	win.ShowAndRun()
	appl.Quit()

	wg.Wait()
}
