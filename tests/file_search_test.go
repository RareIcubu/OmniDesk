package main_test

import (
	"testing"
	"file_manager/fileops"
	"os"
	"path/filepath"
)

func setupTestFiles(t *testing.T, baseDir string) {
	// Tworzymy folder testowy i przykładowe pliki
	err := os.MkdirAll(baseDir, 0755)
	if err != nil {
		t.Fatalf("Nie udało się utworzyć folderu testowego: %v", err)
	}

	files := []string{"file1.txt", "file2.log", "main.go", "file3.TXT"}
	for _, file := range files {
		path := filepath.Join(baseDir, file)
		f, err := os.Create(path)
		if err != nil {
			t.Fatalf("Nie udało się utworzyć pliku testowego: %v", err)
		}
		f.Close()
	}
}

func cleanupTestFiles(t *testing.T, baseDir string) {
	// Usuwamy folder testowy i jego zawartość
	err := os.RemoveAll(baseDir)
	if err != nil {
		t.Errorf("Nie udało się usunąć folderu testowego: %v", err)
	}
}

func TestSearchFile(t *testing.T) {
	// Przygotowanie środowiska testowego
	testDir := "test_files"
	setupTestFiles(t, testDir)
	defer cleanupTestFiles(t, testDir)

	// Test: wyszukiwanie pliku "file1.txt" (case-sensitive)
	var items []fileops.FileItem
	fileops.SearchFile(nil, testDir, "file1.txt", true, &items)

	if len(items) == 0 {
		t.Errorf("Nie znaleziono pliku 'file1.txt'")
	} else if items[0].Name != "file1.txt" {
		t.Errorf("Oczekiwano 'file1.txt', otrzymano '%s'", items[0].Name)
	}

	// Test: wyszukiwanie z ignorowaniem wielkości liter
	items = []fileops.FileItem{} // Resetujemy listę wyników
	fileops.SearchFile(nil, testDir, "FILE3.txt", false, &items)

	if len(items) != 1 {
		t.Errorf("Oczekiwano 1 wyniku dla 'FILE3.txt', otrzymano %d", len(items))
	} else if items[0].Name != "file3.TXT" {
		t.Errorf("Oczekiwano 'file3.TXT', otrzymano '%s'", items[0].Name)
	}

	// Test: wyszukiwanie z regexem
	items = []fileops.FileItem{} // Resetujemy listę wyników
	fileops.SearchFile(nil, testDir, "file.*", false, &items)

	if len(items) != 3 {
		t.Errorf("Oczekiwano 3 wyników dla 'file.*', otrzymano %d", len(items))
	}
}


