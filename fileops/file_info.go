package fileops

import (
	"fmt"
	"os"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// GetFileInfo zwraca informacje o pliku pod podaną ścieżką
func GetFileInfo(filePath string) (os.FileInfo, error) {
	return os.Stat(filePath)
}

// FileInfoDialog tworzy i wyświetla okno dialogowe z informacjami o pliku
func FileInfoDialog(win fyne.Window, filePath string) {
	fileInfo, err := GetFileInfo(filePath)
	if err != nil {
		// Wyświetlamy błąd w dialogu
		dialog.ShowError(err, win)
		return
	}

	// Jeśli plik istnieje, tworzymy layout z jego informacjami
	if fileInfo != nil {
		// Wybieramy ikonę w zależności od typu pliku
		icon := theme.DocumentIcon()
		if fileInfo.IsDir() {
			icon = theme.FolderIcon()
		}

		// Tworzymy duży obraz dla ikony
		iconImage := canvas.NewImageFromResource(icon)
		iconImage.SetMinSize(fyne.NewSize(100, 100)) // Ustawiamy minimalny rozmiar ikony
		iconImage.FillMode = canvas.ImageFillContain

		// Layout dla informacji o pliku
		infoContainer := container.NewVBox(
			iconImage, // Duża ikona
			widget.NewLabel(fmt.Sprintf("Nazwa: %s", fileInfo.Name())),    // Nazwa pliku
			widget.NewLabel(fmt.Sprintf("Rozmiar: %d bajtów", fileInfo.Size())), // Rozmiar pliku
			widget.NewLabel(fmt.Sprintf("Ostatnia modyfikacja: %s", fileInfo.ModTime().String())), // Czas modyfikacji
			widget.NewLabel(fmt.Sprintf("Uprawnienia: %s", fileInfo.Mode().String())),            // Uprawnienia pliku
		)

		// Wyświetlamy okno dialogowe z informacjami
		dialog.ShowCustom("Informacje o pliku", "Zamknij", infoContainer, win)
	}
}

