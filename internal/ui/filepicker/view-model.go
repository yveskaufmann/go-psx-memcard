package filepicker

import (
	"com.yv35.memcard/internal/memcard"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
)

type ViewModel struct {
	FilePath   binding.String
	OnChanged  func(filePath string)
	filePicker FilePickerService
}

func NewViewModel(filePicker FilePickerService) *ViewModel {
	return &ViewModel{
		FilePath:   binding.NewString(),
		filePicker: filePicker,
	}
}

func (v *ViewModel) SetFilePath(filePath string) {
	v.FilePath.Set(filePath)
	if v.OnChanged != nil {
		v.OnChanged(filePath)
	}
}

func (v *ViewModel) GetFilePath() string {
	val, err := v.FilePath.Get()
	if err != nil {
		return ""
	}
	return val
}

func (v *ViewModel) PickFileCommand() {
	currentPath, _ := v.FilePath.Get()
	selectedPath, err := v.filePicker.PickFile(currentPath)

	if err != nil {
		dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
		return
	}

	if selectedPath != "" {
		v.SetFilePath(selectedPath)
		if v.OnChanged != nil {
			v.OnChanged(selectedPath)
		}
	}

}

func (v *ViewModel) CreateNewFileCommand() {
	currentPath, _ := v.FilePath.Get()
	selectedPath, err := v.filePicker.SaveFile(currentPath)

	if err != nil {
		dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
		return
	}

	if selectedPath != "" {
		// Create a new formatted memory card and write it to the selected path
		card := memcard.NewFormattedMemoryCard()
		if err := card.Write(selectedPath); err != nil {
			dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		// Set the file path and trigger the callback to load the new card
		v.SetFilePath(selectedPath)
		if v.OnChanged != nil {
			v.OnChanged(selectedPath)
		}
	}
}
