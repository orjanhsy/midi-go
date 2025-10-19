package frontend

import (
	"image/color"
	"log"
	"midi/backend"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateSettingsScreen(
	prefs fyne.Preferences,
	bottomBar *fyne.Container,
) *fyne.Container {
	onNoteButtonClicked := func(val bool) {
		log.Printf("[SETTINGS] showNote set to %t\n", val)
		prefs.SetBool("showNote", val)
	}

	noteOnButton := widget.NewCheck("Vis noter", onNoteButtonClicked)
	noteOnButton.SetChecked(prefs.BoolWithFallback("showNote", true))

	noteColors := backend.LoadNoteColors(prefs)

	rows := []fyne.CanvasObject{}
	for note, col := range noteColors {
		row := createColorRow(note, col, func(note string, c color.RGBA) {
			noteColors[note] = c
			backend.SaveNoteColors(prefs, noteColors)

			log.Printf("Updated the color of %s\n", note)
		})
		rows = append(rows, row)
	}

	noteColorContainer := container.NewVBox(rows...)
	scrollRows := container.NewVScroll(noteColorContainer)

	screen := container.NewBorder(noteOnButton, bottomBar, nil, nil, scrollRows)
	return screen
}

func createColorRow(note string, col color.RGBA, onChange func(string, color.RGBA)) fyne.CanvasObject {
	label := widget.NewLabel(note + " →")
	label.TextStyle = fyne.TextStyle{Monospace: true}
	labelContainer := container.NewStack(label)
	labelContainer.Resize(fyne.NewSize(40, label.MinSize().Height))

	rEntry := widget.NewEntry()
	gEntry := widget.NewEntry()
	bEntry := widget.NewEntry()

	for _, e := range []*widget.Entry{rEntry, gEntry, bEntry} {
		e.MultiLine = false
		e.Scroll = fyne.ScrollNone
		e.Wrapping = fyne.TextWrapOff
	}

	rEntry.SetText(strconv.Itoa(int(col.R)))
	gEntry.SetText(strconv.Itoa(int(col.G)))
	bEntry.SetText(strconv.Itoa(int(col.B)))

	// preview rectangle
	preview := canvas.NewRectangle(col)
	preview.SetMinSize(fyne.NewSize(20, 20))

	clamp := func(v int) int {
		if v < 0 {
			return 0
		}
		if v > 255 {
			return 255
		}
		return v
	}

	update := func() {
		r, _ := strconv.Atoi(rEntry.Text)
		g, _ := strconv.Atoi(gEntry.Text)
		b, _ := strconv.Atoi(bEntry.Text)

		r = clamp(r)
		g = clamp(g)
		b = clamp(b)

		// update text only if changed
		if rEntry.Text != strconv.Itoa(r) {
			rEntry.SetText(strconv.Itoa(r))
		}
		if gEntry.Text != strconv.Itoa(g) {
			gEntry.SetText(strconv.Itoa(g))
		}
		if bEntry.Text != strconv.Itoa(b) {
			bEntry.SetText(strconv.Itoa(b))
		}

		// update map
		onChange(note, color.RGBA{uint8(r), uint8(g), uint8(b), 255})

		// update preview rectangle
		preview.FillColor = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		preview.Refresh()
	}

	rEntry.OnChanged = func(_ string) { update() }
	gEntry.OnChanged = func(_ string) { update() }
	bEntry.OnChanged = func(_ string) { update() }

	rContainer := container.NewHBox(rEntry, layout.NewSpacer())
	gContainer := container.NewHBox(gEntry, layout.NewSpacer())
	bContainer := container.NewHBox(bEntry, layout.NewSpacer())

	rContainer.Resize(fyne.NewSize(48, rEntry.MinSize().Height))
	gContainer.Resize(fyne.NewSize(48, gEntry.MinSize().Height))
	bContainer.Resize(fyne.NewSize(48, bEntry.MinSize().Height))

	row := container.NewHBox(
		label,
		widget.NewLabel("R:"), rEntry,
		widget.NewLabel("G:"), gEntry,
		widget.NewLabel("B:"), bEntry,
		preview,
	)

	return container.NewPadded(row)
}
