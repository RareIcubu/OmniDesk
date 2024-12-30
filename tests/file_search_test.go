package main

import (
    "testing"
    "file_manager/fileops"
)

func TestSearchFile(t *testing.T) {
    // Test case 1
    // Search for a file that exists in the directory
    // Expected: File found
    file := "file1.txt"
    dir := "test_files"
found, err := fileops.SearchFile(file, dir)
