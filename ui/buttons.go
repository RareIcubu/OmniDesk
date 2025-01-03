package ui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"path/filepath"
	"file_manager/fileops"
)

// CreateGlobalButtons creates global buttons that interact with the active tab.
func CreateGlobalButtons(myWindow fyne.Window, tabs *container.DocTabs, tabStates map[*container.TabItem]*TabState) *fyne.Container {
	backButton := widget.NewButton("Wróć", func() {
		currentTab := tabs.Selected()
		if currentTab == nil {
			dialog.ShowInformation("Błąd", "Brak aktywnej karty", myWindow)
			return
		}

		state := tabStates[currentTab]
		parentPath := filepath.Dir(*state.CurrentPath)
		if parentPath != *state.CurrentPath {
			*state.CurrentPath = parentPath
			UpdateTabContent(myWindow, state)
		} else {
			dialog.ShowInformation("Info", "Już jesteś w katalogu głównym!", myWindow)
		}
	})

	enterButton := widget.NewButton("Wejdź", func() {
		currentTab := tabs.Selected()
		if currentTab == nil {
			dialog.ShowInformation("Błąd", "Brak aktywnej karty", myWindow)
			return
		}

		state := tabStates[currentTab]
		if *state.SelectedIndex >= 0 && (*state.Items)[*state.SelectedIndex].IsDir {
			*state.CurrentPath = (*state.Items)[*state.SelectedIndex].Path
			UpdateTabContent(myWindow, state)
		} else {
			dialog.ShowInformation("Błąd", "Wybierz folder, aby do niego wejść!", myWindow)
		}
	})

	editButton := widget.NewButton("Edytuj", func() {
		currentTab := tabs.Selected()
		if currentTab == nil {
			dialog.ShowInformation("Błąd", "Brak aktywnej karty", myWindow)
			return
		}

		state := tabStates[currentTab]
		if *state.SelectedIndex >= 0 && !(*state.Items)[*state.SelectedIndex].IsDir {
			file := (*state.Items)[*state.SelectedIndex]

			// Add a new tab for editing
			editContent := CreateEditTabContent(myWindow, file.Path, tabs)
			editTab := container.NewTabItem("Edycja: "+file.Name, editContent)
			tabs.Append(editTab)
			tabs.Select(editTab)
		} else {
			dialog.ShowInformation("Błąd", "Wybierz plik, aby go edytować!", myWindow)
		}
	})

	sortButton := widget.NewButton("Sortuj", func() {
		currentTab := tabs.Selected()
		if currentTab == nil {
			dialog.ShowError(errors.New("Brak aktywnej karty"), myWindow)
			fileops.Logger.Println("Brak aktywnej karty")
			return
		}

		state := tabStates[currentTab]

		// Sprawdzanie, czy lista nie jest pusta
		if state.Items == nil || len(*state.Items) == 0 {
			dialog.ShowInformation("Błąd", "Nie ma nic do posortowania!", myWindow)
			return
		}

		// Sortowanie elementów
		fileops.SortItems(state.Items)

		// Reset indeksu zaznaczenia
		*state.SelectedIndex = -1

		// Aktualizacja listy w GUI
		state.List.Refresh()
	})

	openFolderButton := widget.NewButton("Otwórz folder", func() {
		fileops.OpenFolderDialog(myWindow, nil, nil)
	})

	infoButton := widget.NewButton("Info", func() {
		currentTab := tabs.Selected()
		if currentTab == nil {
			dialog.ShowInformation("Błąd", "Brak aktywnej karty", myWindow)
			return
		}

		state := tabStates[currentTab]
		if *state.SelectedIndex >= 0 {
			fileops.FileInfoDialog(myWindow, (*state.Items)[*state.SelectedIndex].Path)
		} else {
			dialog.ShowInformation("Błąd", "Wybierz element, aby zobaczyć szczegóły!", myWindow)
		}
	})

	searchButton := widget.NewButton("Szukaj", func() {
		currentTab := tabs.Selected()
		if currentTab == nil {
			dialog.ShowInformation("Błąd", "Brak aktywnej karty", myWindow)
			return
		}

		state := tabStates[currentTab]
		state.SearchContainer.Show()
	})

	return container.NewHBox(
		backButton,
		enterButton,
		editButton,
		sortButton,
		openFolderButton,
		infoButton,
		searchButton,
	)
}

