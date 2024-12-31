package ui

import (
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"file_manager/fileops"
)

// CreateButtons tworzy przyciski na dole okna.
func CreateButtons(
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
			UpdateList(myWindow, currentPath, items, list, showCurrentPathLabel)
		} else {
			dialog.ShowInformation("Błąd", "Wybierz folder, aby do niego wejść!", myWindow)
		}
	})
	backButton := widget.NewButton("Wróć", func() {
		parentPath := filepath.Dir(*currentPath)
		if parentPath != *currentPath {
			*currentPath = parentPath
			UpdateList(myWindow, currentPath, items, list, showCurrentPathLabel)
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
