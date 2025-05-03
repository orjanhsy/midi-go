package state

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	"midi/backend"
	"midi/clrconv"
)

type ListenerState struct {
	color string
}

func (ls *ListenerState) SetColor(color string) {
	ls.color = color
}

func (ls *ListenerState) GetColor() string {
	return ls.color
}

func (ls *ListenerState) SetNoteHandler(rect *canvas.Rectangle, lab *canvas.Text) {
	handler := func(newCol string, newNote string) {
		ls.SetColor(newCol)
		rectCol, err := clrconv.GetRGBAFromReadableColor(ls.color)
		if err != nil {
			log.Println("Failed to convert color")
		}

		fyne.Do(
			func() {
				rect.FillColor = rectCol
				rect.Refresh()
				lab.Text = newNote
				lab.Refresh()
			},
		)
	}
	backend.SetNoteRecievedHandler(
		handler,
	)
}
