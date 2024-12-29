package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	// Tworzymy aplikację i główne okno.
	myApp := app.New()
	myWindow := newMainWindow(myApp)

	// Uruchamiamy aplikację.
	myWindow.ShowAndRun()
}


