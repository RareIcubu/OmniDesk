package ui

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
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
	currentPath := "." // Bieżąca ścieżka katalogu
	selectedIndex := -1

	// Tworzenie listy z ikonami
	list := widget.NewList(
		func() int {
			return len(items)
		},
		func() fyne.CanvasObject {
			icon := widget.NewIcon(nil)
			label := widget.NewLabel("")
			return container.NewHBox(icon, label)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			item := items[id]
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

	// Pole tekstowe do wyszukiwania
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Wyszukaj plik...")

	// Funkcja aktualizująca zawartość listy
	updateList := func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		items = []fileops.FileItem{}

		for _, file := range files {
			items = append(items, fileops.FileItem{
				Name:  file.Name(),
				Path:  filepath.Join(path, file.Name()),
				IsDir: file.IsDir(),
			})
		}
		list.Refresh()
	}

	// Obsługa wyszukiwania
	searchButton := widget.NewButton("Szukaj", func() {
		search := searchEntry.Text
		if search == "" {
			dialog.ShowInformation("Błąd", "Wpisz nazwę pliku", myWindow)
			return
		}

		var results []fileops.FileItem
		fileops.SearchFile(myWindow, currentPath, search, false, &results)

		if len(results) == 0 {
			dialog.ShowInformation("Wynik", "Nie znaleziono żadnych plików", myWindow)
			return
		}

		items = results
		list.Refresh()
	})

	// Obsługa wyboru elementu z listy
	list.OnSelected = func(id widget.ListItemID) {
		selectedIndex = id
		fmt.Printf("Wybrano element: %s\n", items[id].Name)
	}

	// Przywracamy funkcję `OpenFolderDialog` z fileops/file_opener.go
	openFolder := func() {
		fileops.OpenFolderDialog(myWindow, &items, list)
	}

	// Przycisk "Otwórz folder"
	folderButton := widget.NewButton("Otwórz folder", openFolder)

	// Przycisk "Wejdź do folderu"
	enterButton := widget.NewButton("Wejdź", func() {
		if selectedIndex >= 0 && items[selectedIndex].IsDir {
			currentPath = items[selectedIndex].Path
			updateList(currentPath)
		} else {
			dialog.ShowInformation("Błąd", "Wybierz folder, aby do niego wejść!", myWindow)
		}
	})

	// Przycisk "Wróć"
	backButton := widget.NewButton("Wróć", func() {
		if currentPath != "/" {
			currentPath = filepath.Dir(currentPath)
			updateList(currentPath)
		} else {
			dialog.ShowInformation("Info", "Już jesteś w katalogu głównym!", myWindow)
		}
	})

	// Przycisk "Info"
	infoButton := widget.NewButton("Info", func() {
		if selectedIndex >= 0 {
			fileops.FileInfoDialog(myWindow, items[selectedIndex].Path)
		} else {
			dialog.ShowInformation("Błąd", "Wybierz element, aby zobaczyć szczegóły!", myWindow)
		}
	})

	// Przycisk "Sortuj"
	sortButton := widget.NewButton("Sortuj", func() {
		fileops.SortItems(&items)
		list.Refresh()
	})

	// Menu "Plik"
	fileMenu := fyne.NewMenu("Plik",
		fyne.NewMenuItem("Otwórz folder", openFolder),
	)
	myWindow.SetMainMenu(fyne.NewMainMenu(fileMenu))

	// Layout główny
	layout := container.NewBorder(
		container.NewAdaptiveGrid(2,searchEntry, searchButton),
		container.NewHBox(backButton, enterButton, infoButton, folderButton, sortButton),
		nil,
		nil,
		list,
	)
	myWindow.SetContent(layout)

	// Inicjalizacja zawartości listy
	updateList(currentPath)

	return myWindow
}

