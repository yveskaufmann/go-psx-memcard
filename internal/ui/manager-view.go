package ui

import (
	"com.yv35.memcard/internal/memcard"
	"com.yv35.memcard/internal/ui/blocks"
	"com.yv35.memcard/internal/ui/filepicker"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ManagerWindowView struct {
	model     *ManagerWindowViewModel
	container *fyne.Container
}

func NewManagerWindowView(window fyne.Window) *ManagerWindowView {
	model := NewManagerWindowViewModel(window)
	view := &ManagerWindowView{
		model: model,
	}

	rootLayout := container.NewVBox()

	leftMemoryCardView := blocks.NewContainer(memcard.MemoryCardLeft, model.blocksLeft, model.selection)
	leftMemoryCardView.SetOnBlockSelected(model.HandleBlockSelectionChanged)

	leftMemoryCardFilePicker := filepicker.NewFilePicker(&window)
	leftMemoryCardFilePicker.SetOnChanged(func(filePath string) {
		model.LoadMemoryCardImage(filePath, memcard.MemoryCardLeft)
	})

	leftMemcardContainer := container.NewVBox(
		widget.NewLabel("Card 1"),
		leftMemoryCardFilePicker,
		leftMemoryCardView,
	)

	rightMemoryCardView := blocks.NewContainer(memcard.MemoryCardRight, model.blocksRight, model.selection)
	rightMemoryCardView.SetOnBlockSelected(model.HandleBlockSelectionChanged)

	rightMemoryCardFilePicker := filepicker.NewFilePicker(&window)
	rightMemoryCardFilePicker.SetOnChanged(func(filePath string) {
		model.LoadMemoryCardImage(filePath, memcard.MemoryCardRight)
	})

	rightMemoryCardContainer := container.NewVBox(
		widget.NewLabel("Card 2"),
		rightMemoryCardFilePicker,
		rightMemoryCardView,
	)

	buttons := container.NewVBox()
	btnCopy := widget.NewButton("Copy", func() {
		if err := model.CopyCommand(model.SelectedCard(), model.SelectedBlockIndex()); err != nil {
			dialog.ShowError(err, window)
		}
	})

	btnDelete := widget.NewButton("Delete", func() {
		if err := model.DeleteCommand(model.SelectedCard(), model.SelectedBlockIndex()); err != nil {
			dialog.ShowError(err, window)
		}
	})

	buttons.Add(layout.NewSpacer())
	buttons.Add(btnCopy)
	buttons.Add(btnDelete)
	buttons.Add(layout.NewSpacer())

	labelSelectedSaveGame := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{
		Bold: true,
	})

	labelSelectedSaveGame.Bind(model.selectedSaveGameTitle)

	rootLayout.Add(container.NewBorder(
		nil,
		labelSelectedSaveGame,
		leftMemcardContainer,
		rightMemoryCardContainer,
		buttons,
	))

	view.container = rootLayout

	return view
}

func (v *ManagerWindowView) Container() *fyne.Container {
	return v.container
}
