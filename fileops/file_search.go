package fileops

import (
	"os"
	"path/filepath"
	"regexp"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
    "fmt"
)

// RegexCompile kompiluje wzorzec wyszukiwania z uwzględnieniem wielkości liter
func RegexCompile(search string, caseSensitive bool) (*regexp.Regexp, error) {
	if search == "" {
		return nil, fmt.Errorf("Wzorzec wyszukiwania nie może być pusty")
	}

	if caseSensitive {
		return regexp.Compile(search)
	}
	return regexp.Compile("(?i)" + search) // Ignorowanie wielkości liter
}

// RecursiveSearchFile rekurencyjnie przeszukuje foldery według wzorca regex
func RecursiveSearchFile(myWindow fyne.Window, path string, items *[]FileItem, regex *regexp.Regexp) error {
	err := filepath.WalkDir(path, func(filePath string, d os.DirEntry, err error) error {
		if err != nil {
			dialog.ShowError(err, myWindow)
			return err // Zwracamy błąd, aby przerwać działanie
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

	if err != nil {
		return err
	}

	return nil
}

// SearchFile - funkcja wyszukująca pliki pasujące do wzorca
func SearchFile(myWindow fyne.Window, path string, search string, caseSensitive bool, items *[]FileItem) {
	// Kompilujemy wyrażenie regularne
	regex, err := RegexCompile(search, caseSensitive)
	if err != nil {
		dialog.ShowError(err, myWindow)
		return
	}

	// Wywołujemy funkcję rekurencyjnego wyszukiwania
	err = RecursiveSearchFile(myWindow, path, items, regex)
	if err != nil {
		dialog.ShowError(err, myWindow)
	}
}

