package frontend

import (
	"image/color"
	"midi/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func CreateListenerScreen(
	ls state.ListenerState,
	bottomBar *fyne.Container,
	pref fyne.Preferences,
) *fyne.Container {
	col := ls.GetColor()

	rect := canvas.NewRectangle(col)
	noteLab := canvas.NewText("", color.Black)
	noteLab.TextSize = 128
	noteLab.TextStyle.Bold = true

	ls.SetNoteHandler(rect, noteLab, pref)

	center := container.NewStack(
		rect,
		container.NewCenter(noteLab),
	)

	listenerScreen := container.NewBorder(nil, bottomBar, nil, nil, center)
	return listenerScreen
}
