package main

import (
    "testing"
    "os"
    "fmt"
    "file_manager/fileops"
)

func TestGetFileInfo(t *testing.T) {
    file, err := os.Create("test.txt")
    if err != nil {
        fmt.Println(err)
    }
    os.WriteFile("test.txt", []byte("Hello World"), 0644)
    file.Close()
    fileInfo, err := fileops.GetFileInfo("test.txt")
    if err != nil {
        t.Errorf("Error getting file info")
    }
    if fileInfo.Size() == 0 {
        t.Errorf("File size is 0")
    }
    os.Remove("test.txt")
}
