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
	// Tworzenie listy z ikonami
	list := widget.NewList(
		func() int {
			return len(items)
		},
		func() fyne.CanvasObject {
			// Ikona i etykieta dla każdego elementu
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

currentPath := "." // Bieżąca ścieżka katalogu
	selectedIndex := -1

	// Funkcja aktualizująca zawartość listy na podstawie bieżącej ścieżki
	updateList := func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		// Czyścimy poprzednią zawartość
		items = []fileops.FileItem{}

		// Dodajemy pliki i foldery do listy
		for _, file := range files {
			items = append(items, fileops.FileItem{
				Name:  file.Name(),
				Path:  filepath.Join(path, file.Name()),
				IsDir: file.IsDir(),
			})
		}

		// Odświeżamy listę po zmianie zawartości
		list.Refresh()
	}
	// Obsługa wyboru elementu z listy
	list.OnSelected = func(id widget.ListItemID) {
		selectedIndex = id
		fmt.Printf("Wybrano element: %s\n", items[id].Name)
	}

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

	// Layout główny: przyciski na dole + lista
	layout := container.NewBorder(
		widget.NewLabel("File Manager"),
		container.NewHBox(backButton, enterButton, infoButton, sortButton),
		nil,
		nil,
		list,
	)
	myWindow.SetContent(layout)

	// Inicjalizacja zawartości listy
	updateList(currentPath)

	return myWindow
}

