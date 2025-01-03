package ui

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
)

// CreateCurrentPathLabel tworzy etykietę wyświetlającą bieżącą ścieżkę.
func CreateCurrentPathLabel(currentPath string) *widget.Label {
	return widget.NewLabel(fmt.Sprintf("Ścieżka: %s", currentPath))
}

