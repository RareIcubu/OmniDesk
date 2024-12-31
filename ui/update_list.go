package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"file_manager/fileops"
)

// UpdateList aktualizuje zawartość listy i etykiety bieżącej ścieżki.
func UpdateList(
	myWindow fyne.Window,
	currentPath *string,
	items *[]fileops.FileItem,
	list *widget.List,
	showCurrentPathLabel *widget.Label,
) {
	err := fileops.UpdateList(*currentPath, items)
	if err != nil {
		dialog.ShowError(err, myWindow)
		return
	}

	showCurrentPathLabel.SetText(fmt.Sprintf("Ścieżka: %s", *currentPath))
	list.Refresh()
}
