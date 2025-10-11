package filepicker

import (
	"os"
	"path"

	"com.yvka.memcard/pkg/ui/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

type FilePickerService interface {
	// PickFile opens a file dialog and returns the selected file path.
	// If the initialPath is not empty, it will be used as the starting directory.
	PickFile(initialPath string) (string, error)
}

func DetermineInitialLocation(currentPath string) string {
	if currentPath != "" && path.Ext(currentPath) != "" {
		return path.Dir(currentPath)
	}

	if currentPath != "" {
		return currentPath
	}

	homeDir, err := os.UserHomeDir()
	if err == nil {
		return homeDir
	}

	return ""
}

type FyneFilePickerService struct {
	window *fyne.Window
}

func NewFyneFilePickerService(window *fyne.Window) *FyneFilePickerService {
	return &FyneFilePickerService{
		window: window,
	}
}

func (s *FyneFilePickerService) PickFile(initialPath string) (string, error) {

	fc := make(chan string)
	fe := make(chan error)

	window := *s.window
	if window == nil {
		window = utils.GetPrimaryWindow()
	}

	uri := storage.NewFileURI(DetermineInitialLocation(initialPath))
	lister, err := storage.ListerForURI(uri)
	if err != nil {
		dialog.ShowError(err, window)
		return "", err
	}

	fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			fe <- err
			return
		}

		if reader == nil {
			return
		}

		defer reader.Close()

		path := reader.URI().Path()
		if path != "" {
			fc <- path
		}
	}, window)

	fileDialog.SetLocation(lister)
	fileDialog.Show()

	select {
	case filePath := <-fc:
		return filePath, nil
	case err := <-fe:
		return "", err
	}
}
