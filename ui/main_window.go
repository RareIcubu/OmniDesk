package ui

import (
	"fmt"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"file_manager/fileops"
)

// NewMainWindow tworzy główne okno aplikacji.
func NewMainWindow(a fyne.App) fyne.Window {
	myWindow := a.NewWindow("File Manager")
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.CenterOnScreen()

	// Inicjalizacja danych
	currentPath, _ := filepath.Abs(".")
	var items []fileops.FileItem
	selectedIndex := -1

	// Elementy UI
	showCurrentPathLabel := widget.NewLabel(fmt.Sprintf("Ścieżka: %s", currentPath))
	list := fileops.CreateList(&items, &selectedIndex)
	buttons := createButtons(myWindow, &currentPath, &items, list, &selectedIndex, showCurrentPathLabel)

	// Tworzymy menu
	fileMenu := createFileMenu(myWindow, &items, list)
	myWindow.SetMainMenu(fyne.NewMainMenu(fileMenu))

	// Layout główny
	layout := container.NewBorder(
		container.NewVBox(showCurrentPathLabel),
		buttons,
		nil,
		nil,
		list,
	)
	myWindow.SetContent(layout)

	// Inicjalizacja listy
	updateList(myWindow, &currentPath, &items, list, showCurrentPathLabel)

	return myWindow
}

// createButtons tworzy przyciski na dole okna.
func createButtons(
	myWindow fyne.Window,
	currentPath *string,
	items *[]fileops.FileItem,
	list *widget.List,
	selectedIndex *int,
	showCurrentPathLabel *widget.Label,
) *fyne.Container {
	backButton := widget.NewButton("Wróć", func() {
		parentPath := filepath.Dir(*currentPath)
		if parentPath != *currentPath {
			*currentPath = parentPath
			updateList(myWindow, currentPath, items, list, showCurrentPathLabel)
		} else {
			dialog.ShowInformation("Info", "Już jesteś w katalogu głównym!", myWindow)
		}
	})

	enterButton := widget.NewButton("Wejdź", func() {
		if *selectedIndex >= 0 && (*items)[*selectedIndex].IsDir {
			*currentPath = (*items)[*selectedIndex].Path
			updateList(myWindow, currentPath, items, list, showCurrentPathLabel)
		} else {
			dialog.ShowInformation("Błąd", "Wybierz folder, aby do niego wejść!", myWindow)
		}
	})

	infoButton := widget.NewButton("Info", func() {
		if *selectedIndex >= 0 {
			fileops.FileInfoDialog(myWindow, (*items)[*selectedIndex].Path)
		} else {
			dialog.ShowInformation("Błąd", "Wybierz element, aby zobaczyć szczegóły!", myWindow)
		}
	})

	sortButton := widget.NewButton("Sortuj", func() {
		fileops.SortItems(items)
		list.Refresh()
	})

	folderButton := widget.NewButton("Otwórz folder", func() {
		fileops.OpenFolderDialog(myWindow, items, list)
	})

	searchButton := widget.NewButton("Szukaj", func() {
		showSearchDialog(myWindow, *currentPath, items, list)
	})

	return container.NewHBox(backButton, enterButton, infoButton, folderButton, sortButton, searchButton)
}

// createFileMenu tworzy menu "Plik".
func createFileMenu(myWindow fyne.Window, items *[]fileops.FileItem, list *widget.List) *fyne.Menu {
	return fyne.NewMenu("Plik",
		fyne.NewMenuItem("Otwórz folder", func() {
			fileops.OpenFolderDialog(myWindow, items, list)
		}),
	)
}

// updateList aktualizuje zawartość listy i etykiety bieżącej ścieżki.
func updateList(
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

// showSearchDialog otwiera okno dialogowe do wyszukiwania plików.
func showSearchDialog(myWindow fyne.Window, currentPath string, items *[]fileops.FileItem, list *widget.List) {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Wpisz nazwę pliku...")
	searchButton := widget.NewButton("Szukaj", func() {
		search := searchEntry.Text
		if search == "" {
			dialog.ShowInformation("Błąd", "Wpisz nazwę pliku!", myWindow)
			return
		}

		var results []fileops.FileItem
		fileops.SearchFile(myWindow, currentPath, search, false, &results)

		if len(results) == 0 {
			dialog.ShowInformation("Wynik", "Nie znaleziono żadnych plików!", myWindow)
			return
		}

		*items = results
		list.Refresh()
	})

	content := container.NewVBox(searchEntry, searchButton)
	dialog.ShowCustom("Wyszukiwanie", "Zamknij", content, myWindow)
}

