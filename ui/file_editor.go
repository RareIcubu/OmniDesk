package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"log"
)

// OpenFileEditor creates a new tab to edit the selected file.
func OpenFileEditor(filePath string, tabs *container.DocTabs) {
	// Load file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to open file: %s", err)
		return
	}

	// Create a TextGrid for editing
	textGrid := widget.NewTextGrid()
	textGrid.SetText(string(content))

	// Save button
	saveButton := widget.NewButton("Save", func() {
		newContent := textGrid.Text()
		err := os.WriteFile(filePath, []byte(newContent), 0644)
		if err != nil {
			log.Printf("Failed to save file: %s", err)
		}
	})

	// Add a new tab for file editing
	editorTab := container.NewTabItem("Edit: "+filePath, container.NewBorder(nil, saveButton, nil, nil, textGrid))
	tabs.Append(editorTab)
	tabs.Select(editorTab) // Automatically switch to the new tab
}
