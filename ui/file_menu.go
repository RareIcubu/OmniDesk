package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"file_manager/fileops"
    "path/filepath"
    "fyne.io/fyne/v2/dialog"
)

// CreateMainMenu creates the main menu for the application.
func CreateMainMenu(myWindow fyne.Window, tabs *container.DocTabs, tabStates map[*container.TabItem]*TabState) *fyne.MainMenu {
	return fyne.NewMainMenu(
		fyne.NewMenu("Plik",
			fyne.NewMenuItem("Otwórz folder", func() {
				dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
					if err != nil {
						dialog.ShowError(err, myWindow)
						return
					}
					if uri != nil {
						currentPath := uri.Path()
						newItems := []fileops.FileItem{}
						newSelectedIndex := -1
						newShowPathLabel := CreateCurrentPathLabel(currentPath)

						newTabContent, newTabState := CreateTabContent(myWindow, &currentPath, &newItems, &newSelectedIndex, newShowPathLabel, tabs)
						newTab := container.NewTabItem(filepath.Base(currentPath), newTabContent)

						tabs.Append(newTab)
						tabs.Select(newTab)
						tabStates[newTab] = newTabState
					}
				}, myWindow)
			}),
		),
	)
}

