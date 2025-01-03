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
	// Back Button
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

	// Enter Folder Button
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

	// Edit File Button
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
			editTab := container.NewTabItem(file.Name, editContent)
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
	    if state == nil {
		    dialog.ShowError(errors.New("Nie można znaleźć stanu dla aktywnej karty"), myWindow)
		    fileops.Logger.Println("Nie można znaleźć stanu dla aktywnej karty")
		    return
	    }

	    // Sortowanie elementów
	    fileops.SortItems(state.Items)

	    // Aktualizacja wyświetlenia po sortowaniu
	    UpdateTabContent(myWindow, state)
    })
	// Open Folder Button
	openFolderButton := widget.NewButton("Otwórz folder", func() {
		 dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri != nil {
				currentPath := uri.Path()
				newItems := []fileops.FileItem{}
				newSelectedIndex := -1
				newShowPathLabel := CreateCurrentPathLabel(currentPath)

				newTabContent, newTabState := CreateTabContent(myWindow, &currentPath, &newItems, &newSelectedIndex, newShowPathLabel, tabs)
				newTab := container.NewTabItem(filepath.Base(currentPath), newTabContent)

				tabs.Append(newTab)
				tabStates[newTab] = newTabState
				tabs.Select(newTab)
                if err != nil {
			        dialog.ShowError(err, myWindow)
		        }
			}
		}, myWindow)
		
	})

	// Info Button
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

	// Search Button
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

