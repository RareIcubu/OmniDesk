package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"file_manager/fileops"
    "path/filepath"

)

// NewMainWindow tworzy główne okno aplikacji.
func NewMainWindow(a fyne.App) fyne.Window {
	myWindow := a.NewWindow("File Manager")
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.CenterOnScreen()

	currentPath, _ := filepath.Abs(".")
	var items []fileops.FileItem
	selectedIndex := -1

	// Elementy UI
	showCurrentPathLabel := createCurrentPathLabel(currentPath)
	list := CreateFileList(&items, &selectedIndex)

	// Deklaracja kontenera wyszukiwania
	var searchContainer *fyne.Container
	searchContainer = CreateSearchContainer(myWindow, &currentPath, &items, list, searchContainer)

	// Przyciski
	buttons := CreateButtons(myWindow, &currentPath, &items, list, &selectedIndex, showCurrentPathLabel, searchContainer)

	// Menu
	fileMenu := CreateFileMenu(myWindow, &items, list)
	myWindow.SetMainMenu(fyne.NewMainMenu(fileMenu))

	// Layout
	layout := container.NewBorder(
		container.NewVBox(showCurrentPathLabel, searchContainer),
		buttons,
		nil,
		nil,
		list,
	)
	myWindow.SetContent(layout)

	// Inicjalizacja listy
	UpdateList(myWindow, &currentPath, &items, list, showCurrentPathLabel)

	return myWindow
}

