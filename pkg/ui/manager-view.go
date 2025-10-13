package ui

import (
	"com.yvka.memcard/pkg/memcard"
	"com.yvka.memcard/pkg/ui/blocks"
	"com.yvka.memcard/pkg/ui/filepicker"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ManagerWindowView struct {
	model     *ManagerWindowViewModel
	container *fyne.Container
}

func NewManagerWindowView(window fyne.Window) *ManagerWindowView {
	view := &ManagerWindowView{
		model: &ManagerWindowViewModel{
			window: window,
		},
	}

	rootLayout := container.NewVBox()

	rootLayout.Add(widget.NewLabel("Select a memory card file to begin."))

	leftMemoryCardView := blocks.NewBlockContainer(memcard.MemoryCardLeft)
	leftMemoryCardFilePicker := filepicker.NewFilePicker(&window)
	leftMemoryCardFilePicker.SetOnChanged(func(filePath string) {
		view.model.LoadMemoryCardImage(filePath, memcard.MemoryCardLeft)
	})

	leftMemcardContainer := container.NewVBox(
		widget.NewLabel("Card 1"),
		leftMemoryCardFilePicker,
		leftMemoryCardView,
	)

	rightMemoryCardView := blocks.NewBlockContainer(memcard.MemoryCardRight)
	rightMemoryCardFilePicker := filepicker.NewFilePicker(&window)
	rightMemoryCardFilePicker.SetOnChanged(func(filePath string) {
		view.model.LoadMemoryCardImage(filePath, memcard.MemoryCardRight)
	})

	rightMemoryCardContainer := container.NewVBox(
		widget.NewLabel("Card 2"),
		rightMemoryCardFilePicker,
		rightMemoryCardView,
	)

	buttons := container.NewVBox()
	btnCopy := widget.NewButton("Copy", func() {
		view.model.CopyCommand(memcard.MemoryCardLeft, view.model.SelectedBlockIndex())
	})

	btnCopyAll := widget.NewButton("Copy All", func() {
		view.model.CopyAllCommand(memcard.MemoryCardLeft)
	})

	btnDelete := widget.NewButton("Delete", func() {
		view.model.DeleteCommand(memcard.MemoryCardRight, view.model.SelectedBlockIndex())
	})

	buttons.Add(layout.NewSpacer())
	buttons.Add(btnCopy)
	buttons.Add(btnCopyAll)
	buttons.Add(btnDelete)
	buttons.Add(layout.NewSpacer())

	rootLayout.Add(container.NewHBox(
		leftMemcardContainer,
		layout.NewSpacer(),
		buttons,
		layout.NewSpacer(),
		rightMemoryCardContainer,
	))

	view.container = rootLayout

	return view
}

func (v *ManagerWindowView) Container() *fyne.Container {
	return v.container
}
