package ui

import (
	"image/color"

	"com.yv35.memcard/internal/memcard"
	"com.yv35.memcard/internal/ui/blocks"
	"com.yv35.memcard/internal/ui/blockstats"
	"com.yv35.memcard/internal/ui/filepicker"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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

	leftMemoryCardView := blocks.NewContainer(memcard.MemoryCardLeft, model.blocksLeft, model.selection)
	leftMemoryCardView.SetOnBlockSelected(model.HandleBlockSelectionChanged)

	leftMemoryCardFilePicker := filepicker.NewFilePicker(&window)
	leftMemoryCardFilePicker.SetOnChanged(func(filePath string) {
		model.LoadMemoryCardImage(filePath, memcard.MemoryCardLeft)
	})

	// Create styled card header
	card1Header := createCardHeader("Card 1")

	leftMemcardContainer := container.NewVBox(
		card1Header,
		leftMemoryCardFilePicker,
		leftMemoryCardView,
	)

	rightMemoryCardView := blocks.NewContainer(memcard.MemoryCardRight, model.blocksRight, model.selection)
	rightMemoryCardView.SetOnBlockSelected(model.HandleBlockSelectionChanged)

	rightMemoryCardFilePicker := filepicker.NewFilePicker(&window)
	rightMemoryCardFilePicker.SetOnChanged(func(filePath string) {
		model.LoadMemoryCardImage(filePath, memcard.MemoryCardRight)
	})

	// Create styled card header
	card2Header := createCardHeader("Card 2")

	rightMemoryCardContainer := container.NewVBox(
		card2Header,
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

	// Create container for the selected save game label (will be populated dynamically)
	selectedSaveGameContainer := container.NewVBox()

	// Create label for selected save game with placeholder support
	// We'll recreate the label with different styles based on selection state
	var labelSelectedSaveGame *widget.Label

	// Update label text and style based on selection state
	updateSelectedGameLabel := func() {
		text, _ := model.selectedSaveGameTitle.Get()
		if text == "" {
			// Show placeholder text when no block is selected (grayed out, not bold)
			labelSelectedSaveGame = widget.NewLabelWithStyle("No block selected", fyne.TextAlignCenter, fyne.TextStyle{
				Bold: false,
			})
			labelSelectedSaveGame.Importance = widget.LowImportance // Grayed out placeholder appearance
		} else {
			// Show actual save game title with bold styling
			labelSelectedSaveGame = widget.NewLabelWithStyle(text, fyne.TextAlignCenter, fyne.TextStyle{
				Bold: true,
			})
			labelSelectedSaveGame.Importance = widget.MediumImportance // Normal appearance
		}
		// Update the container content
		selectedSaveGameContainer.RemoveAll()
		selectedSaveGameContainer.Add(labelSelectedSaveGame)
		selectedSaveGameContainer.Refresh()
	}

	// Listen to title changes
	model.selectedSaveGameTitle.AddListener(binding.NewDataListener(updateSelectedGameLabel))

	// Initial update
	updateSelectedGameLabel()

	// Create block statistics view
	blockStatsView := blockstats.NewMemoryCardBlockStatsView(
		func() (total, used, free int) {
			return model.GetBlockStatistics(memcard.MemoryCardLeft)
		},
		func() (total, used, free int) {
			return model.GetBlockStatistics(memcard.MemoryCardRight)
		},
	)

	// Listen to block binding changes to update statistics
	model.blocksLeft.AddListener(binding.NewDataListener(func() {
		blockStatsView.UpdateStatistics()
	}))
	model.blocksRight.AddListener(binding.NewDataListener(func() {
		blockStatsView.UpdateStatistics()
	}))

	// Initial update
	blockStatsView.UpdateStatistics()

	// Create footer container that will stay at the bottom
	footerContainer := container.NewVBox(
		selectedSaveGameContainer,
		blockStatsView.Container(),
	)

	// Use Border layout as root to ensure footer stays at bottom
	// The center area (buttons) will expand to fill available space, pushing footer down
	view.container = container.NewBorder(
		nil,
		footerContainer, // Footer always at bottom
		leftMemcardContainer,
		rightMemoryCardContainer,
		buttons, // Center area expands to fill space
	)

	return view
}

func (v *ManagerWindowView) Container() *fyne.Container {
	return v.container
}

// createCardHeader creates a visually appealing header for a memory card section.
// It includes a styled background, border, and formatted text.
func createCardHeader(title string) *fyne.Container {
	// Create label with bold, larger text
	label := widget.NewLabelWithStyle(title, fyne.TextAlignCenter, fyne.TextStyle{
		Bold: true,
	})

	// Create background rectangle with subtle color
	backgroundRect := canvas.NewRectangle(color.RGBA{
		R: 230,
		G: 230,
		B: 235,
		A: 255,
	})

	// Create border rectangle
	borderRect := canvas.NewRectangle(color.Transparent)
	borderRect.StrokeColor = color.RGBA{R: 200, G: 200, B: 200, A: 255}
	borderRect.StrokeWidth = 1
	borderRect.FillColor = color.Transparent

	// Create container with padding for the label
	paddedLabel := container.NewPadded(label)

	// Stack background, border, and label
	headerContainer := container.NewStack(
		backgroundRect,
		container.NewBorder(
			nil,
			borderRect, // Bottom border
			nil,
			nil,
			paddedLabel,
		),
	)

	return headerContainer
}
