package ui

import (
	"fmt"

	"com.yvka.memcard/pkg/memcard"
	"com.yvka.memcard/pkg/ui/filepicker"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Start() error {

	a := app.New()
	w := a.NewWindow("PSX Memory Card Manager")

	rootLayout := container.NewVBox()

	rootLayout.Add(widget.NewLabel("Select a memory card file to begin."))

	leftMemoryCardView := NewBlockContainer(0)
	leftmemCardFilePicker := filepicker.NewFilePicker(&w)
	leftmemCardFilePicker.SetOnChanged(func(filePath string) {
		card, err := memcard.Open(filePath)
		if err != nil {
			dialog.ShowError(err, w)
		}

		fmt.Printf("Loaded memory card: %+v\n", card)
		blocks, err := card.ListBlocks()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		for idx, block := range blocks {
			leftMemoryCardView.SetBlockItem(idx, block)

		}

	})

	leftMemcardContainer := container.NewVBox(
		widget.NewLabel("Card 1"),
		leftmemCardFilePicker,
		leftMemoryCardView,
	)

	rightMemoryCardView := NewBlockContainer(1)
	rightmemCardFilePicker := filepicker.NewFilePicker(&w)

	rightMemcardContainer := container.NewVBox(
		widget.NewLabel("Card 2"),
		rightmemCardFilePicker,
		rightMemoryCardView,
	)

	buttons := container.NewVBox()

	btnCopy := widget.NewButton("Copy", func() {

	})

	btnCopyAll := widget.NewButton("Copy All", func() {

	})

	btnDelete := widget.NewButton("Delete", func() {

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
		rightMemcardContainer,
	))

	w.Resize(fyne.NewSize(1024, 768))
	w.SetContent(rootLayout)
	w.ShowAndRun()

	return nil
}
