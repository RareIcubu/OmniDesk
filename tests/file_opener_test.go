package main

import (
    "testing"
    "file_manager"
)

func TestSortItes(t *testing.T) {
    items := []fileItem{
        {name: "file1.txt", isDir: false},
        {name: "folder1", isDir: true},
        {name: "file2.txt", isDir: false},
        {name: "folder2", isDir: true},
    }
    expected := []fileItem{
        {name: "folder1", isDir: true},
        {name: "folder2", isDir: true},
        {name: "file1.txt", isDir: false},
        {name: "file2.txt", isDir: false},
    }
    sortItems(items)
    
    for i, item := range items {
        if item != expected[i] {
            t.Errorf("Expected %v, got %v", expected[i], item)
        }
    }
}
