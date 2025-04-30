package state

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"

	"midi/backend"
	"midi/clrconv"
)

type ListenerState struct {
	color binding.String
}

func (ls *ListenerState) SetColor(color string) {
	ls.color = binding.BindString(&color)
}

func (ls *ListenerState) GetColor() binding.String {
	return ls.color
}

func (ls *ListenerState) SetNoteHandler(rect *canvas.Rectangle) {
	handler := func(newCol string) {
		ls.SetColor(newCol)
		rectCol, err := clrconv.GetRGBAFromReadableColor(ls.color)
		if err != nil {
			log.Println("Failed to convert color")
		}

		fyne.Do(
			func() {
				rect.FillColor = rectCol
				rect.Refresh()
			},
		)
	}
	backend.SetNoteRecievedHandler(
		handler,
	)
}
