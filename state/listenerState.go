package state

import (
	"image/color"
	"midi/backend"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type ListenerState struct {
	color color.RGBA
}

func (ls *ListenerState) Init(color color.RGBA) {
	ls.SetColor(color)
}

func (ls *ListenerState) SetColor(color color.RGBA) {
	ls.color = color
}

func (ls *ListenerState) GetColor() color.RGBA {
	return ls.color
}

func (ls *ListenerState) SetNoteHandler(
	rect *canvas.Rectangle,
	lab *canvas.Text,
	pref fyne.Preferences,
) {
	handler := func(newCol color.RGBA, newNote string) {
		ls.SetColor(newCol)
		fyne.Do(
			func() {
				rect.FillColor = newCol
				rect.Refresh()
				lab.Text = newNote

				if pref.BoolWithFallback("showNote", true) {
					lab.Show()
				} else {
					lab.Hide()
				}

				lab.Refresh()
			},
		)
	}
	backend.SetNoteRecievedHandler(
		handler,
	)
}
