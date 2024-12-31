package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"file_manager/fileops"
)

// CreateFileMenu tworzy menu "Plik".
func CreateFileMenu(myWindow fyne.Window, items *[]fileops.FileItem, list *widget.List) *fyne.Menu {
	return fyne.NewMenu("Plik",
		fyne.NewMenuItem("Otwórz folder", func() {
			fileops.OpenFolderDialog(myWindow, items, list)
		}),
	)
}
