package main

import (
	"fmt"
	"image/color"
	"sync"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"midi/backend"
	"midi/clrconv"
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

	colorChangeHandler := func(newCol string) {
		boundClrLab.Set(newCol)
		rectClr, err := clrconv.GetRGBAFromReadableColor(boundClrLab)
		if err != nil {
			return
		}
		rect.FillColor = rectClr
		rect.Refresh()
	}

	go func() {
		defer wg.Done()
		backend.ListenForMidiInput(boundClrLab, colorChangeHandler)
	}()
	fmt.Printf("Backend up\n")

	win.SetContent(
		container.NewBorder(clrLab, nil, nil, nil, rect),
	)

	win.ShowAndRun()
	appl.Quit()

	wg.Wait()
}
