package ui

import (
	"fmt"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
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
	list := createFileList(&items, &selectedIndex)

	// Deklaracja kontenera wyszukiwania, aby był widoczny globalnie w tej funkcji
	var searchContainer *fyne.Container

	// Inicjalizacja kontenera wyszukiwania
	searchContainer = createSearchContainer(myWindow, &currentPath, &items, list, searchContainer)

	// Przyciski
	buttons := createButtons(myWindow, &currentPath, &items, list, &selectedIndex, showCurrentPathLabel, searchContainer)

	// Tworzymy menu
	fileMenu := createFileMenu(myWindow, &items, list)
	myWindow.SetMainMenu(fyne.NewMainMenu(fileMenu))

	// Układ główny
	layout := container.NewBorder(
		container.NewVBox(showCurrentPathLabel, searchContainer),
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

// createFileList tworzy listę plików i folderów z obsługą `OnSelected`.
func createFileList(items *[]fileops.FileItem, selectedIndex *int) *widget.List {
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
			*selectedIndex = -1 // Resetujemy index, jeśli nic nie jest wybrane
		}
	}

	return list
}

// createButtons tworzy przyciski na dole okna.
func createButtons(
	myWindow fyne.Window,
	currentPath *string,
	items *[]fileops.FileItem,
	list *widget.List,
	selectedIndex *int,
	showCurrentPathLabel *widget.Label,
	searchContainer *fyne.Container,
) *fyne.Container {
	folderButton := widget.NewButton("Otwórz folder", func() {
		fileops.OpenFolderDialog(myWindow, items, list)
	})
	enterButton := widget.NewButton("Wejdź", func() {
		if *selectedIndex >= 0 && (*items)[*selectedIndex].IsDir {
			*currentPath = (*items)[*selectedIndex].Path
			updateList(myWindow, currentPath, items, list, showCurrentPathLabel)
		} else {
			dialog.ShowInformation("Błąd", "Wybierz folder, aby do niego wejść!", myWindow)
		}
	})
	backButton := widget.NewButton("Wróć", func() {
		parentPath := filepath.Dir(*currentPath)
		if parentPath != *currentPath {
			*currentPath = parentPath
			updateList(myWindow, currentPath, items, list, showCurrentPathLabel)
		} else {
			dialog.ShowInformation("Info", "Już jesteś w katalogu głównym!", myWindow)
		}
	})
	infoButton := widget.NewButton("Info", func() {
		if *selectedIndex >= 0 && *selectedIndex < len(*items) {
			fileops.FileInfoDialog(myWindow, (*items)[*selectedIndex].Path)
		} else {
			dialog.ShowInformation("Błąd", "Wybierz element, aby zobaczyć szczegóły!", myWindow)
		}
	})
	sortButton := widget.NewButton("Sortuj", func() {
		fileops.SortItems(items)
		list.Refresh()
	})
	searchButton := widget.NewButton("Szukaj", func() {
		searchContainer.Show()
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

// createSearchContainer tworzy wysuwany pasek wyszukiwania.
func createSearchContainer(
	myWindow fyne.Window,
	currentPath *string,
	items *[]fileops.FileItem,
	list *widget.List,
	searchContainer *fyne.Container,
) *fyne.Container {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Wpisz nazwę pliku...")

	searchButton := widget.NewButton("Szukaj", func() {
		search := searchEntry.Text
		if search == "" {
			dialog.ShowInformation("Błąd", "Wpisz nazwę pliku!", myWindow)
			return
		}

		var results []fileops.FileItem
		fileops.SearchFile(myWindow, *currentPath, search, false, &results)

		if len(results) == 0 {
			dialog.ShowInformation("Wynik", "Nie znaleziono żadnych plików!", myWindow)
			return
		}

		*items = results
		list.Refresh()
	})

	closeButton := widget.NewButton("Zamknij", func() {
		searchContainer.Hide()  // Ukrywa cały kontener
		searchEntry.SetText("") // Czyści tekst
	})

	searchContainer = container.NewVBox(searchEntry, container.NewHBox(searchButton, closeButton))
	searchContainer.Hide() // Ukrywamy kontener na początku
	return searchContainer
}

