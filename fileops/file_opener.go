package fileops

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/container"
	"sort"
)

// FileItem przechowuje nazwę pliku/podfolderu i informację, czy to folder.
type FileItem struct {
	Name  string
    Path string
	IsDir bool
}

// CreateFileMenu tworzy menu z opcją do wybrania folderu i wyświetlenia w liście.
func CreateFileMenu(win fyne.Window, items *[]FileItem, list *widget.List) *fyne.Menu {
	return fyne.NewMenu("File Manager",
		fyne.NewMenuItem("List Folder", func() {
			OpenFolderDialog(win, items, list)
		}),
	)
}

// OpenFolderDialog wyświetla dialog do wyboru folderu i aktualizuje listę.
func OpenFolderDialog(win fyne.Window, items *[]FileItem, list *widget.List) {
	dialog.ShowFolderOpen(
		func(listable fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if listable != nil {
				UpdateList(listable, items, list, win)
			}
		},
		win,
	)
}

// UpdateList odczytuje zawartość folderu, aktualizuje slice `items` i odświeża listę.
func UpdateList(listable fyne.ListableURI, items *[]FileItem, list *widget.List, win fyne.Window) {
	children, err := listable.List()
	if err != nil {
		dialog.ShowError(err, win)
		return
	}

	*items = (*items)[:0]

	for _, child := range children {
		_, listerErr := storage.ListerForURI(child)
		isDir := (listerErr == nil)

		*items = append(*items, FileItem{
			Name:  child.Name(),
			IsDir: isDir,
            Path: child.Path(),
		})
	}

	list.Refresh()
}

// CreateList tworzy widget.List dla wyświetlania plików i folderów.
func CreateList(items *[]FileItem) *widget.List {
	return widget.NewList(
		func() int {
			return len(*items)
		},
		func() fyne.CanvasObject {
			icon := widget.NewIcon(theme.FileIcon())
			label := widget.NewLabel("")
			return container.NewHBox(icon, label)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			container := obj.(*fyne.Container)
			icon := container.Objects[0].(*widget.Icon)
			label := container.Objects[1].(*widget.Label)

			fi := (*items)[id]
			if fi.IsDir {
				icon.SetResource(theme.FolderIcon())
			} else {
				icon.SetResource(theme.FileIcon())
			}
			label.SetText(fi.Name)
		},
	)
}

// SortItems sortuje listę plików i folderów: najpierw foldery, potem pliki.
func SortItems(items *[]FileItem) {
	sort.Slice(*items, func(i, j int) bool {
		if (*items)[i].IsDir && !(*items)[j].IsDir {
			return true
		}
		if !(*items)[i].IsDir && (*items)[j].IsDir {
			return false
		}
		return (*items)[i].Name < (*items)[j].Name
	})
}

