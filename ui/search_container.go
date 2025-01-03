
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"file_manager/fileops"
)

// CreateSearchContainer creates the search container.
func CreateSearchContainer(
	myWindow fyne.Window,
	currentPath *string,
	items *[]fileops.FileItem,
	list *widget.List,
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
		searchEntry.SetText("")
	})

	searchContainer := container.NewVBox(searchEntry, container.NewHBox(searchButton, closeButton))
	searchContainer.Hide()
	return searchContainer
}

