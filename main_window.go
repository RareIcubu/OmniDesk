package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)
func newMainWindow(a fyne.App) fyne.Window {
	myWindow := a.NewWindow("File Manager")

	// Ustawiamy rozmiar i pozycję okna.
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.CenterOnScreen()

	// Lista elementów w folderze.
	var items []fileItem

	// Tworzymy listę plików/podfolderów.
	list := createList(&items)

	// Tworzymy menu "File Manager".
	fileMenu := createFileMenu(myWindow, &items, list)
	myWindow.SetMainMenu(fyne.NewMainMenu(fileMenu))

    folderButton := widget.NewButton("Open Folder", func() {
        openFolderDialog(myWindow,&items, list)
    })
    sortButton := widget.NewButton("Sort", func() {
        sortItems(&items)
        list.Refresh()
    })
    // Layout główny: etykieta + lista.
    layout := container.NewBorder(
        widget.NewLabel("File Manager"),
        container.NewHBox(folderButton, sortButton),
        nil,
        nil,
        list,
    )
	myWindow.SetContent(layout)

	return myWindow
}


