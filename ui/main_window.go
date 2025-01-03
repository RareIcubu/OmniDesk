
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"path/filepath"
	"file_manager/fileops"
)

// NewMainWindow creates the main application window.
func NewMainWindow(a fyne.App) fyne.Window {
	myWindow := a.NewWindow("File Manager")
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.CenterOnScreen()

	// Initialize state
	currentPath, _ := filepath.Abs(".")
	var items []fileops.FileItem
	selectedIndex := -1

	// Create UI elements
	showCurrentPathLabel := CreateCurrentPathLabel(currentPath)
	tabs := container.NewDocTabs()
	tabStates := make(map[*container.TabItem]*TabState) // Map to store tab states

	// Create global buttons
	buttons := CreateGlobalButtons(myWindow, tabs, tabStates)

	// Add default tab
	defaultTabContent, defaultTabState := CreateTabContent(myWindow, &currentPath, &items, &selectedIndex, showCurrentPathLabel, tabs)
	defaultTab := container.NewTabItem("File Manager", defaultTabContent)
	tabs.Append(defaultTab)
	tabStates[defaultTab] = defaultTabState
	tabs.SetTabLocation(container.TabLocationTop)

	// Create the MainMenu
	mainMenu := CreateMainMenu(myWindow, tabs, tabStates) 
	myWindow.SetMainMenu(mainMenu)

	// Layout
	mainLayout := container.NewBorder(
		container.NewVBox(buttons),
		nil,
		nil,
		nil,
		tabs,
	)

	myWindow.SetContent(mainLayout)
	return myWindow
}

