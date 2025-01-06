# OmniDesk

OmniDesk to wielofunkcyjny menedżer plików stworzony przy użyciu frameworka [Fyne](https://fyne.io). Jest to narzędzie zaprojektowane w celu zwiększenia produktywności użytkowników, oferujące zaawansowane funkcje zarządzania plikami, edycji i przeszukiwania katalogów w jednym miejscu.

---

## 📋 Cele projektu

Głównym celem OmniDesk jest:

1. **Zwiększenie produktywności**:
   - Integracja różnych funkcji zarządzania plikami w jednej aplikacji.
   - Intuicyjny i responsywny interfejs użytkownika.

2. **Uniwersalność**:
   - Możliwość pracy z wieloma folderami i plikami jednocześnie dzięki zakładkom.

3. **Rozbudowane funkcje edycyjne**:
   - Stylizacja edytowanych plików.
   - Automatyczne śledzenie zmian w plikach.

4. **Łatwość obsługi**:
   - Prosty interfejs, który nie wymaga zaawansowanej wiedzy technicznej.
   - Wbudowane mechanizmy logowania, które pomagają w debugowaniu.

---

## ✨ Kluczowe funkcje

### 1. **Przeglądanie plików i folderów**
   - Wyświetlanie zawartości folderów.
   - Obsługa wielu folderów w zakładkach (DocTabs).
   - Informacje o plikach:
     - Rozmiar
     - Typ (plik/katalog)
     - Ostatnia modyfikacja

### 2. **Edycja plików**
   - Edytor tekstu oparty na `TextGrid`:

### 3. **Sortowanie plików**
   - Sortowanie plików alfabetycznie.
   - Foldery wyświetlane na początku listy.

### 4. **Zaawansowane wyszukiwanie**
   - Wyszukiwanie plików w bieżącym folderze oraz rekurencyjnie w podfolderach.
   - Obsługa wyrażeń regularnych (regex).
   - Ignorowanie wielkości liter (opcjonalnie).

### 5. **Obsługa zakładek**
   - Możliwość pracy z wieloma folderami i plikami jednocześnie w osobnych zakładkach.
   - Edytor dla każdego pliku w osobnej zakładce.

### 6. **Logowanie zdarzeń**
   - Rejestrowanie błędów i zdarzeń w pliku logów (`file_manager.log`).
   - Logowanie takich zdarzeń jak otwieranie folderów, zapisywanie plików czy błędy.

---

## 🛠️ Wymagania systemowe

- Go 1.20 lub nowszy
- System operacyjny:
  - Windows
  - macOS
  - Linux

---

## 🚀 Instalacja i uruchomienie

### Klonowanie repozytorium

```bash
git clone https://github.com/username/omnidesk.git
cd omnidesk
```

### Instalacja zależności

OmniDesk opiera się na frameworku [Fyne](https://fyne.io). Aby upewnić się, że wszystkie zależności są zainstalowane, wykonaj:

```bash
go mod tidy
```

### Uruchomienie aplikacji

Aby uruchomić OmniDesk, użyj następującego polecenia:

```bash
go run cmd/main.go
```

Lub skompiluj aplikację i uruchom binarny plik wykonywalny:

```bash
go build -o omnidesk cmd/main.go
./omnidesk
```

---

## 📂 Struktura projektu

```
.
├── cmd
│   └── main.go                # Główna funkcja aplikacji
├── fileops
│   ├── file_info.go           # Operacje związane z informacjami o plikach
│   ├── file_opener.go         # Operacje otwierania folderów i plików
│   ├── file_search.go         # Mechanizmy wyszukiwania plików
│   ├── logger.go              # Obsługa logów
├── tests
│   ├── file_info_test.go      # Testy dla file_info.go
│   ├── file_opener_test.go    # Testy dla file_opener.go
│   ├── file_search_test.go    # Testy dla file_search.go
├── ui
│   ├── buttons.go             # Definicje przycisków interfejsu
│   ├── current_path_label.go  # Etykiety z bieżącą ścieżką
│   ├── edit_tab.go            # Zakładki edycji plików
│   ├── file_editor.go         # Edytor plików
│   ├── file_list.go           # Lista plików i folderów
│   ├── file_menu.go           # Menu aplikacji
│   ├── main_window.go         # Główne okno aplikacji
│   ├── search_container.go    # Kontener wyszukiwania
│   ├── tabs_content.go        # Zarządzanie zawartością zakładek
│   ├── update_list.go         # Aktualizacja listy plików
├── go.mod
├── go.sum
└── file_manager.log           # Plik logów aplikacji
```

---

## 📖 Przykłady użycia

### Otwieranie folderów
1. Kliknij przycisk "Otwórz folder".
2. Wybierz folder z wyświetlonego dialogu.
3. Zawartość folderu pojawi się w nowej zakładce.

### Edycja plików
1. Wybierz plik z listy.
2. Kliknij przycisk "Edytuj".
3. Plik otworzy się w edytorze tekstowym z numerami linii.
4. Edytuj treść pliku i kliknij "Zapisz", aby zapisać zmiany.

### Wyszukiwanie plików
1. Wprowadź nazwę pliku w pasku wyszukiwania.
2. Kliknij przycisk "Szukaj".
3. Wyniki wyszukiwania zostaną wyświetlone na liście.

---

## 🔧 Rozwój i testowanie

### Uruchamianie testów

Testy znajdują się w folderze `tests`. Aby uruchomić wszystkie testy, wykonaj:

```bash
go test ./...
```

### Dodawanie nowych funkcji

Nowe funkcje powinny być dodawane w odpowiednich modułach (`fileops` lub `ui`) i pokryte testami w folderze `tests`.

---

## 🖋️ Licencja

OmniDesk jest dostępny na licencji GNU/GPL. Szczegóły znajdziesz w pliku `LICENSE`.

---

## 👥 Autorzy

- **Jakub Jasiński** - Główny autor projektu.

---

Dziękujemy za korzystanie z OmniDesk! 😊

