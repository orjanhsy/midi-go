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
	prefs fyne.Preferences,
) {
	handler := func(newNote string) {
		noteColors := backend.LoadNoteColors(prefs)
		fyne.Do(
			func() {
				rect.FillColor = noteColors[newNote]
				rect.Refresh()
				switch newNote {
				case "Db":
					lab.Text = "C#"
				case "Gb":
					lab.Text = "F#"
				default:
					lab.Text = newNote
				}

				if prefs.BoolWithFallback("showNote", true) {
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
