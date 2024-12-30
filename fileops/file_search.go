package fileops

import (
	"os"
	"path/filepath"
	"regexp"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// SearchFile - funkcja wyszukująca pliki pasujące do wzorca
func SearchFile(myWindow fyne.Window, path string, search string, caseSensitive bool, items *[]FileItem) {
	// Kompilujemy wyrażenie regularne raz na początku
	var regex *regexp.Regexp
	var err error
	if caseSensitive {
		regex, err = regexp.Compile(search)
	} else {
		regex, err = regexp.Compile("(?i)" + search) // Ignorowanie wielkości liter
	}
	if err != nil {
		dialog.ShowError(err, myWindow)
		return
	}

	// Rekurencyjne przeszukiwanie folderów
	err = filepath.WalkDir(path, func(filePath string, d os.DirEntry, err error) error {
		if err != nil {
			// Zwracamy błąd tylko raz dla całej funkcji
			dialog.ShowError(err, myWindow)
			return nil
		}

		// Pomijamy foldery (ale nadal wchodzimy do nich)
		if d.IsDir() {
			return nil
		}

		// Sprawdzamy, czy nazwa pliku pasuje do wzorca
		if regex.MatchString(d.Name()) {
			*items = append(*items, FileItem{
				Name:  d.Name(),
				Path:  filePath, // Dodajemy pełną ścieżkę pliku
				IsDir: false,
			})
		}
		return nil
	})

	// Wyświetlamy błąd, jeśli coś poszło nie tak z WalkDir
	if err != nil {
		dialog.ShowError(err, myWindow)
	}
}

