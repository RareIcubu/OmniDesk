package main

import (
	"fyne.io/fyne/v2/app"
    "file_manager/ui"
)

func main() {
	// Tworzymy aplikację i główne okno.
	myApp := app.New()
	myWindow := ui.NewMainWindow(myApp)

	// Uruchamiamy aplikację.
	myWindow.ShowAndRun()
}


