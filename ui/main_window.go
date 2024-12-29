package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"file_manager/fileops"
)

func NewMainWindow(a fyne.App) fyne.Window {
	myWindow := a.NewWindow("File Manager")

	// Ustawiamy rozmiar i pozycję okna.
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.CenterOnScreen()

	// Lista elementów w folderze.
	var items []fileops.FileItem

	// Tworzymy listę plików/podfolderów.
	list := fileops.CreateList(&items)

	// Tworzymy menu "File Manager".
	fileMenu := fileops.CreateFileMenu(myWindow, &items, list)
	myWindow.SetMainMenu(fyne.NewMainMenu(fileMenu))

	// Tworzymy przyciski do otwierania folderu i sortowania.
	folderButton := widget.NewButton("Open Folder", func() {
		fileops.OpenFolderDialog(myWindow, &items, list)
	})
	sortButton := widget.NewButton("Sort", func() {
		fileops.SortItems(&items)
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

