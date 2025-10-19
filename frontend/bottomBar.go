package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Tab string

const (
	TabDevice   Tab = "device"
	TabListener Tab = "listener"
	TabSettings Tab = "settings"
)

func CreateAppBar(
	onDeviceMenuClicked func(),
	onListenerClicked func(),
	onSettingsClicked func(),
) *fyne.Container {
	currentTab := TabDevice // default selected tab

	var updateButtonStyles func() // declare first so callbacks can use it

	deviceButton := widget.NewButton("Enheter", func() {
		currentTab = TabDevice
		onDeviceMenuClicked()
		updateButtonStyles()
	})

	listenerButton := widget.NewButton("Visualiser", func() {
		currentTab = TabListener
		onListenerClicked()
		updateButtonStyles()
	})

	settingsButton := widget.NewButton("Innstillinger", func() {
		currentTab = TabSettings
		onSettingsClicked()
		updateButtonStyles()
	})

	buttons := []*widget.Button{deviceButton, listenerButton, settingsButton}

	// define function after buttons are created
	updateButtonStyles = func() {
		for _, b := range buttons {
			b.Importance = widget.MediumImportance // default
			b.Refresh()
		}

		// highlight current tab
		switch currentTab {
		case TabDevice:
			deviceButton.Importance = widget.HighImportance
			deviceButton.Refresh()
		case TabListener:
			listenerButton.Importance = widget.HighImportance
			listenerButton.Refresh()
		case TabSettings:
			settingsButton.Importance = widget.HighImportance
			settingsButton.Refresh()
		}
	}

	updateButtonStyles() // initial highlight

	bar := container.NewGridWithColumns(3, deviceButton, listenerButton, settingsButton)
	return bar
}

