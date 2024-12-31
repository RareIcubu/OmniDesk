package ui

import (
	"fyne.io/fyne/v2/widget"
)

// createCurrentPathLabel tworzy etykietę wyświetlającą bieżącą ścieżkę.
func createCurrentPathLabel(currentPath string) *widget.Label {
	return widget.NewLabel("Ścieżka: " + currentPath)
}
