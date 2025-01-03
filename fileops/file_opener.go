package fileops

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"sort"
)

// FileItem przechowuje nazwę pliku/podfolderu i informację, czy to folder.
type FileItem struct {
	Name  string
	Path  string
	IsDir bool
}

// OpenFolderDialog wyświetla dialog do wyboru folderu i aktualizuje listę.
func OpenFolderDialog(win fyne.Window, items *[]FileItem, list *widget.List) {
	Logger.Println("Otwarcie dialogu wyboru folderu")
    dialog.ShowFolderOpen(
		func(listable fyne.ListableURI, err error) {
			if err != nil {
                Logger.Println("Błąd podczas otwierania folderu:", err)
				dialog.ShowError(err, win)
				return
			}
			if listable != nil {
                Logger.Printf("Otwarto folder: %s\n", listable.Path())
				UpdateListFromURI(listable, items, list, win)
			}
		},
		win,
	)
}

// UpdateListFromURI aktualizuje listę na podstawie URI.
func UpdateListFromURI(listable fyne.ListableURI, items *[]FileItem, list *widget.List, win fyne.Window) {
	children, err := listable.List()
	if err != nil {
		dialog.ShowError(err, win)
		return
	}

	*items = []FileItem{}
	for _, child := range children {
		isDir := false
		if _, err := storage.ListerForURI(child); err == nil {
			isDir = true
		}

		*items = append(*items, FileItem{
			Name:  child.Name(),
			Path:  child.Path(),
			IsDir: isDir,
		})
	}
	list.Refresh()
}

// UpdateList aktualizuje listę na podstawie ścieżki.
func UpdateList(path string, items *[]FileItem) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	*items = []FileItem{}
	for _, file := range files {
		*items = append(*items, FileItem{
			Name:  file.Name(),
			Path:  filepath.Join(path, file.Name()),
			IsDir: file.IsDir(),
		})
	}
	return nil
}

// CreateList tworzy widget.List dla wyświetlania plików i folderów.
func CreateList(items *[]FileItem, selectedIndex *int) *widget.List {
	return widget.NewList(
		func() int {
			return len(*items)
		},
		func() fyne.CanvasObject {
			icon := widget.NewIcon(nil)
			label := widget.NewLabel("")
			return container.NewHBox(icon, label)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			container := obj.(*fyne.Container)
			icon := container.Objects[0].(*widget.Icon)
			label := container.Objects[1].(*widget.Label)

			item := (*items)[id]
			if item.IsDir {
				icon.SetResource(theme.FolderIcon())
			} else {
				icon.SetResource(theme.FileIcon())
			}
			label.SetText(item.Name)
		},
	)
}

func SortItems(items *[]FileItem) {
	if items == nil || len(*items) == 0 {
		return // Nie ma czego sortować
	}

	sort.Slice(*items, func(i, j int) bool {
		// Najpierw sortuj foldery przed plikami
		if (*items)[i].IsDir && !(*items)[j].IsDir {
			return true
		}
		if !(*items)[i].IsDir && (*items)[j].IsDir {
			return false
		}
		// Następnie sortuj alfabetycznie
		return (*items)[i].Name < (*items)[j].Name
	})
}
