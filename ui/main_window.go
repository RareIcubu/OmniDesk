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
	var selectedIndex int = -1 // Przechowujemy indeks wybranego elementu

	// Tworzymy listę plików/podfolderów.
	list := fileops.CreateList(&items)

	// Obsługa wyboru elementu z listy
	list.OnSelected = func(id widget.ListItemID) {
		if id >= 0 && id < len(items) {
			selectedIndex = id // Zapisujemy indeks wybranego elementu
		}
	}

	// Tworzymy menu "File Manager".
	fileMenu := fileops.CreateFileMenu(myWindow, &items, list)
	myWindow.SetMainMenu(fyne.NewMainMenu(fileMenu))

	// Tworzymy przyciski
	folderButton := widget.NewButton("Open Folder", func() {
		fileops.OpenFolderDialog(myWindow, &items, list)
	})

	sortButton := widget.NewButton("Sort", func() {
		fileops.SortItems(&items)
		list.Refresh()
	})

	fileinfoButton := widget.NewButton("File Info", func() {
		// Wyświetlamy informacje o wybranym elemencie
		if selectedIndex >= 0 && selectedIndex < len(items) {
			fileops.FileInfoDialog(myWindow, items[selectedIndex].Path)
		} else {
			dialog := widget.NewLabel("Wybierz element z listy przed kliknięciem przycisku!")
			myWindow.SetContent(container.NewVBox(dialog, list))
		}
	})

	// Layout główny: etykieta + lista.
	layout := container.NewBorder(
		widget.NewLabel("File Manager"),
		container.NewHBox(folderButton, sortButton, fileinfoButton),
		nil,
		nil,
		list,
	)
	myWindow.SetContent(layout)

	return myWindow
}

