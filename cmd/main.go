package main

import (
	"file_manager/fileops"
	"file_manager/ui"
	"fyne.io/fyne/v2/app"
	"log"
)

func main() {
	// Inicjalizacja loggera
	logFilePath := "file_manager.log" // Ścieżka do pliku logów
	err := fileops.InitLogger(logFilePath)
	if err != nil {
		log.Fatalf("Nie udało się zainicjalizować loggera: %v", err)
	}
	defer fileops.Logger.Println("Zamknięcie aplikacji")

	// Tworzymy aplikację i główne okno
	myApp := app.New()
	myWindow := ui.NewMainWindow(myApp)

	// Uruchamiamy aplikację
	myWindow.ShowAndRun()
}

