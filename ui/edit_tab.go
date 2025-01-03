package ui

import (
	"bufio"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"os"
)

// CreateEditTabContent creates a new tab for editing the file.
func CreateEditTabContent(myWindow fyne.Window, filePath string, tabs *container.DocTabs) fyne.CanvasObject {
	content := widget.NewMultiLineEntry()
	content.Wrapping = fyne.TextWrapWord

	// Read the file content
	file, err := os.Open(filePath)
	if err != nil {
		dialog.ShowError(err, myWindow)
		return widget.NewLabel("Nie udało się załadować pliku.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileContent string
	for scanner.Scan() {
		fileContent += scanner.Text() + "\n"
	}
	content.SetText(fileContent)

	// Save button
	saveButton := widget.NewButton("Zapisz", func() {
		err := os.WriteFile(filePath, []byte(content.Text), os.ModePerm)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Sukces", "Plik zapisany pomyślnie", myWindow)
		}
	})

	// Close tab button
	//closeTabButton := widget.NewButton("Zamknij kartę", func() {
	//    tabs.RemoveIndex(tabs.SelectedIndex())
    //})

	toolbar := container.NewHBox(saveButton)

	return container.NewBorder(toolbar, nil, nil, nil, content)
}

