package frontend

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateSettingsScreen(
	pref fyne.Preferences,
	bottomBar *fyne.Container,
) *fyne.Container {
	onNoteButtonClicked := func(val bool) {
		log.Printf("[SETTINGS] showNote set to %t\n", val)
		pref.SetBool("showNote", val)
	}

	noteOnButton := widget.NewCheck("Vis noter", onNoteButtonClicked)
	noteOnButton.SetChecked(pref.BoolWithFallback("showNote", true))
	screen := container.NewBorder(noteOnButton, bottomBar, nil, nil)
	return screen
}
