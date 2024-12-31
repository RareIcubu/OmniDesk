package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"file_manager/fileops"
)

// CreateFileList tworzy listę plików i folderów z obsługą `OnSelected`.
func CreateFileList(items *[]fileops.FileItem, selectedIndex *int) *widget.List {
	list := widget.NewList(
		func() int {
			return len(*items)
		},
		func() fyne.CanvasObject {
			icon := widget.NewIcon(nil)
			label := widget.NewLabel("")
			return container.NewHBox(icon, label)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			item := (*items)[id]
			container := obj.(*fyne.Container)
			icon := container.Objects[0].(*widget.Icon)
			label := container.Objects[1].(*widget.Label)

			if item.IsDir {
				icon.SetResource(theme.FolderIcon())
			} else {
				icon.SetResource(theme.DocumentIcon())
			}
			label.SetText(item.Name)
		},
	)

	// Obsługa wyboru elementu z listy
	list.OnSelected = func(id widget.ListItemID) {
		if id >= 0 && id < len(*items) {
			*selectedIndex = id
			fmt.Printf("Wybrano element: %s\n", (*items)[id].Name)
		} else {
			*selectedIndex = -1
		}
	}

	return list
}
