
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"file_manager/fileops"
	"path/filepath"
)

// NewMainWindow creates the main window for the application.
func NewMainWindow(a fyne.App) fyne.Window {
	myWindow := a.NewWindow("File Manager")
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.CenterOnScreen()

	// Initialize state
	currentPath, _ := filepath.Abs(".")
	var items []fileops.FileItem
	selectedIndex := -1

	// UI elements
	showCurrentPathLabel := CreateCurrentPathLabel(currentPath)
	list := CreateFileList(&items, &selectedIndex)
	searchContainer := CreateSearchContainer(myWindow, &currentPath, &items, list)

	// Declare tabs first (to use it in buttons)
	tabs := container.NewDocTabs()

	// Create buttons with tabs dependency
	buttons := CreateButtons(myWindow, &currentPath, &items, list, &selectedIndex, showCurrentPathLabel, searchContainer, tabs)

	// Add the File Manager tab
	tabs.Append(container.NewTabItem("File Manager", container.NewBorder(
		container.NewVBox(showCurrentPathLabel, searchContainer),
		buttons,
		nil,
		nil,
		list,
	)))
	tabs.SetTabLocation(container.TabLocationTop)

	// File menu
	fileMenu := CreateFileMenu(myWindow, &items, list)
	myWindow.SetMainMenu(fyne.NewMainMenu(fileMenu))

	myWindow.SetContent(tabs)

	// Initialize list
	UpdateList(myWindow, &currentPath, &items, list, showCurrentPathLabel)

	return myWindow
}

