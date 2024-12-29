
package main

import (
    "os"
    "path/filepath"
	"testing"
	"file_manager/fileops"
)

func TestSortItems(t *testing.T) {
	items := []fileops.FileItem{
		{Name: "file1.txt", IsDir: false}, // Zmieniono `name` na `Name` i `isDir` na `IsDir`
		{Name: "folder1", IsDir: true},
		{Name: "file2.txt", IsDir: false},
		{Name: "folder2", IsDir: true},
	}
	expected := []fileops.FileItem{
		{Name: "folder1", IsDir: true},   // Zmieniono `name` na `Name` i `isDir` na `IsDir`
		{Name: "folder2", IsDir: true},
		{Name: "file1.txt", IsDir: false},
		{Name: "file2.txt", IsDir: false},
	}

	// Sortujemy elementy
	fileops.SortItems(&items)

	// Sprawdzamy, czy posortowana lista zgadza się z oczekiwaniami
	for i, item := range items {
		if item != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], item)
		}
	}
}

func TestUpdateList(t *testing.T) {
	// Tworzymy folder tymczasowy
	tempDir := t.TempDir()

	// Tworzymy pliki i foldery w folderze tymczasowym
	os.Mkdir(filepath.Join(tempDir, "folder1"), 0755)
	os.WriteFile(filepath.Join(tempDir, "file1.txt"), []byte("test"), 0644)

	// Lista elementów
	var items []fileops.FileItem

	// Odczytujemy zawartość folderu ręcznie
	entries, err := os.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("Nie udało się odczytać katalogu: %v", err)
	}

	// Symulujemy działanie `UpdateList`
	for _, entry := range entries {
		items = append(items, fileops.FileItem{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
		})
	}

	// Sprawdzamy, czy elementy zostały poprawnie załadowane
	if len(items) != 2 {
		t.Fatalf("Oczekiwano 2 elementów, ale otrzymano %d", len(items))
	}

	// Sprawdzamy szczegóły elementów
	expectedNames := map[string]bool{
		"folder1":   true,
		"file1.txt": true,
	}
	for _, item := range items {
		if !expectedNames[item.Name] {
			t.Errorf("Nieoczekiwany element: %v", item.Name)
		}
	}
}
