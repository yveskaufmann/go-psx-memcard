package filepicker

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type FilePicker struct {
	widget.BaseWidget

	vm *ViewModel

	// txtFilePath is the text entry widget to display the file path
	txtFilePath *widget.Entry

	// btnBrowse is the button to open the file dialog
	btnBrowse *widget.Button
}

func NewFilePicker(window *fyne.Window) *FilePicker {
	fp := &FilePicker{
		vm: NewViewModel(
			&FyneFilePickerService{window: window},
		),
	}

	// Create the text entry widget
	fp.txtFilePath = widget.NewEntryWithData(fp.vm.FilePath)
	fp.txtFilePath.SetPlaceHolder("Select a file...")

	// Create the browse button
	fp.btnBrowse = widget.NewButton("...", func() {
		go func() {
			fp.vm.PickFileCommand()
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

	c := container.NewBorder(
		nil, nil, nil,
		fp.btnBrowse,
		fp.txtFilePath,
	)
	return widget.NewSimpleRenderer(container.NewStack(c))
}
