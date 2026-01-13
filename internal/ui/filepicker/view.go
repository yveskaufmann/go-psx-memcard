package filepicker

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type FilePicker struct {
	widget.BaseWidget

	vm *ViewModel

	// txtFilePath is the text entry widget to display the file path
	txtFilePath *widget.Entry

	// btnBrowse is the button to open the file dialog
	btnBrowse *widget.Button

	// btnNew is the button to create a new memory card file
	btnNew *widget.Button
}

func NewFilePicker(window *fyne.Window) *FilePicker {
	fp := &FilePicker{
		vm: NewViewModel(
			&FyneFilePickerService{window: window},
		),
	}

	// Create the text entry widget
	fp.txtFilePath = widget.NewEntryWithData(fp.vm.FilePath)
	fp.txtFilePath.SetPlaceHolder("Select a memory card file...")

	// Create the browse button
	fp.btnBrowse = widget.NewButtonWithIcon("", theme.FolderIcon(), func() {
		go func() {
			fp.vm.PickFileCommand()
		}()
	})

	// Create the new file button with document icon
	fp.btnNew = widget.NewButtonWithIcon("", theme.FileIcon(), func() {
		go func() {
			fp.vm.CreateNewFileCommand()
		}()
	})

	fp.ExtendBaseWidget(fp)

	return fp
}

func (fp *FilePicker) SetOnChanged(onChanged func(filePath string)) {
	fp.vm.OnChanged = onChanged
}

func (fp *FilePicker) FilePath() string {
	str, err := fp.vm.FilePath.Get()
	if err != nil {
		return ""
	}
	return str
}

func (fp *FilePicker) CreateRenderer() fyne.WidgetRenderer {
	// Create a horizontal container for the buttons
	buttonContainer := container.NewHBox(fp.btnBrowse, fp.btnNew)

	c := container.NewBorder(
		nil, nil, nil,
		buttonContainer,
		fp.txtFilePath,
	)
	return widget.NewSimpleRenderer(container.NewStack(c))
}
