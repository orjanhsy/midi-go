package frontend

import (
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
	ls.SetNoteHandler(rect)

	listenerScreen := container.NewBorder(nil, bottomBar, nil, nil, rect)
	return listenerScreen
}
