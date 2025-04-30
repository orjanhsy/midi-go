package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateAppBar(
	onDeviceMenuClicked func(),
	onListenerClicked func(),
	onSettingsClicked func(),
) *fyne.Container {
	deviceMenuButton := widget.NewButton("Enheter", onDeviceMenuClicked)
	listenerButton := widget.NewButton("Visualiser", onListenerClicked)
	settingsButton := widget.NewButton("Innstillinger", onSettingsClicked)

	bar := container.NewGridWithColumns(3, deviceMenuButton, listenerButton, settingsButton)
	return bar
}
