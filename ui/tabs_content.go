package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"file_manager/fileops"
)

// TabState holds the state of a single tab.
type TabState struct {
	CurrentPath   *string
	Items         *[]fileops.FileItem
	SelectedIndex *int
	ShowPathLabel *widget.Label
	List          *widget.List
	SearchContainer *fyne.Container
}

// CreateTabContent creates the content of a new tab.
func CreateTabContent(
	myWindow fyne.Window,
	currentPath *string,
	items *[]fileops.FileItem,
	selectedIndex *int,
	showPathLabel *widget.Label,
	tabs *container.DocTabs,
) (fyne.CanvasObject, *TabState) {
	list := CreateFileList(items, selectedIndex)

	searchContainer := CreateSearchContainer(myWindow, currentPath, items, list)
    
    
    showPathLabel.SetText(*currentPath)

	tabState := &TabState{
		CurrentPath:   currentPath,
		Items:         items,
		SelectedIndex: selectedIndex,
		ShowPathLabel: showPathLabel,
		List:          list,
		SearchContainer: searchContainer,
	}

	UpdateTabContent(myWindow, tabState)

	content := container.NewBorder(
		container.NewVBox(showPathLabel, searchContainer),
		nil,
		nil,
		nil,
		list,
	)

	return content, tabState
}

func UpdateTabContent(myWindow fyne.Window, state *TabState) {
	// Aktualizujemy listę plików i folderów
	err := fileops.UpdateList(*state.CurrentPath, state.Items)
	if err != nil {
		dialog.ShowError(err, myWindow)
		return
	}

	// Odświeżamy komponent listy
	state.List.Refresh()
}

