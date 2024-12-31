package fileops

import (
	"log"
	"os"
)

var (
	// Logger to globalny logger używany w aplikacji
	Logger *log.Logger
)

// InitLogger inicjalizuje logger i ustawia plik do zapisywania logów
func InitLogger(logFilePath string) error {
	// Otwieramy plik do zapisu logów
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Konfigurujemy logger
	Logger = log.New(logFile, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

// CloseLogger zamyka plik logów
func CloseLogger(logFile *os.File) {
	logFile.Close()
}

