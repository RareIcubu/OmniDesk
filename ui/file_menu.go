package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"file_manager/fileops"
)

// CreateFileMenu creates the file menu for the application.
func CreateFileMenu(myWindow fyne.Window, items *[]fileops.FileItem, list *widget.List) *fyne.Menu {
	return fyne.NewMenu("Plik",
		fyne.NewMenuItem("Otwórz folder", func() {
			fileops.OpenFolderDialog(myWindow, items, list)
		}),
	)
}

