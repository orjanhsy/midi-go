package frontend

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"midi/clrconv"
	"midi/state"
)

func CreateListenerScreen(ls state.ListenerState, bottomBar *fyne.Container) *fyne.Container {
	col := ls.GetColor()
	rgba, err := clrconv.GetRGBAFromReadableColor(col)
	if err != nil {
		log.Println("Failed to convert color for rectangle in listenerScreen")
	}

	rect := canvas.NewRectangle(rgba)
	noteLab := canvas.NewText("", color.Black)
	noteLab.TextSize = 60
	noteLab.TextStyle.Bold = true

	ls.SetNoteHandler(rect, noteLab)

	center := container.NewStack(
		rect,
		container.NewCenter(noteLab),
	)

	listenerScreen := container.NewBorder(nil, bottomBar, nil, nil, center)
	return listenerScreen
}
