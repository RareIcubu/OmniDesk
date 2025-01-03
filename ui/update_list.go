package ui

import (
    "file_manager/fileops"
	"fmt"
    "os"
    "path/filepath"
)


func UpdateList(path string, items *[]fileops.FileItem) error {
	// Logowanie przed aktualizacją
	fmt.Println("Przed aktualizacją listy:")
	for _, item := range *items {
		fmt.Printf("  - %s (folder: %v)\n", item.Name, item.IsDir)
	}

	// Aktualizujemy zawartość listy
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	var updatedItems []fileops.FileItem
	for _, entry := range entries {
		updatedItems = append(updatedItems, fileops.FileItem{
			Name:  entry.Name(),
			Path:  filepath.Join(path, entry.Name()),
			IsDir: entry.IsDir(),
		})
	}

	// Nadpisujemy wskaźnik
	*items = updatedItems

	// Logowanie po aktualizacji
	fmt.Println("Po aktualizacji listy:")
	for _, item := range *items {
		fmt.Printf("  - %s (folder: %v)\n", item.Name, item.IsDir)
	}

	return nil
}

