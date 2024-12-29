package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/container"
    "sort"
)

// fileItem przechowuje nazwę pliku/podfolderu i informację, czy to folder.
type fileItem struct {
	name  string
	isDir bool
}

// createFileMenu tworzy menu z opcją do wybrania folderu i wyświetlenia w liście.
func createFileMenu(win fyne.Window, items *[]fileItem, list *widget.List) *fyne.Menu {
	return fyne.NewMenu("File Manager",
		fyne.NewMenuItem("List Folder", func() {
			openFolderDialog(win, items, list)
		}),
	)
}

// openFolderDialog wyświetla dialog do wyboru folderu i aktualizuje listę.
func openFolderDialog(win fyne.Window, items *[]fileItem, list *widget.List) {
	dialog.ShowFolderOpen(
		func(listable fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if listable != nil {
				updateList(listable, items, list, win)
			}
		},
		win,
	)
}

// updateList odczytuje zawartość folderu, aktualizuje slice `items` i odświeża listę.
func updateList(listable fyne.ListableURI, items *[]fileItem, list *widget.List, win fyne.Window) {
	// Odczytujemy zawartość folderu.
	children, err := listable.List()
	if err != nil {
		dialog.ShowError(err, win)
		return
	}

	// Czyścimy poprzednią zawartość
	*items = (*items)[:0]

	// Dodajemy nową zawartość
	for _, child := range children {
		_, listerErr := storage.ListerForURI(child)
		isDir := (listerErr == nil)

		*items = append(*items, fileItem{
			name:  child.Name(),
			isDir: isDir,
		})
	}

	// Odświeżamy widok listy po zmianie danych.
	list.Refresh()
}
func createList(items *[]fileItem) *widget.List {
	return widget.NewList(
		// Liczba elementów w liście.
		func() int {
			return len(*items)
		},
		// Jak tworzymy pusty element (ikona + etykieta).
		func() fyne.CanvasObject {
			icon := widget.NewIcon(theme.FileIcon())
			label := widget.NewLabel("")
			return container.NewHBox(icon, label)
		},
		// Jak aktualizujemy element listy.
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			container := obj.(*fyne.Container)
			icon := container.Objects[0].(*widget.Icon)
			label := container.Objects[1].(*widget.Label)

			fi := (*items)[id]
			if fi.isDir {
				icon.SetResource(theme.FolderIcon())
			} else {
				icon.SetResource(theme.FileIcon())
			}
			label.SetText(fi.name)
		},
	)
}
// sortItems sortuje listę plików i folderów: najpierw foldery, potem pliki.
func sortItems(items *[]fileItem) {
	sort.Slice(*items, func(i, j int) bool {
		// Foldery na początku
		if (*items)[i].isDir && !(*items)[j].isDir {
			return true
		}
		if !(*items)[i].isDir && (*items)[j].isDir {
			return false
		}
		// Sortowanie alfabetyczne w obrębie tej samej kategorii
		return (*items)[i].name < (*items)[j].name
	})
}
