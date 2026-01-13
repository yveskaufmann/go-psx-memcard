package blockstats

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// MemoryCardBlockStatsView displays block statistics for two memory cards.
type MemoryCardBlockStatsView struct {
	model     *MemoryCardBlockStatsViewModel
	container *fyne.Container
}

// NewMemoryCardBlockStatsView creates a new view for displaying memory card block statistics.
func NewMemoryCardBlockStatsView(
	getLeftCardStats func() (total, used, free int),
	getRightCardStats func() (total, used, free int),
) *MemoryCardBlockStatsView {
	model := NewMemoryCardBlockStatsViewModel(getLeftCardStats, getRightCardStats)
	view := &MemoryCardBlockStatsView{
		model: model,
	}

	// Create labels for left card statistics
	leftCardTotalLabel := widget.NewLabel("")
	leftCardTotalLabel.Bind(binding.IntToStringWithFormat(model.LeftCardTotal(), "Total: %d"))

	leftCardUsedLabel := widget.NewLabel("")
	leftCardUsedLabel.Bind(binding.IntToStringWithFormat(model.LeftCardUsed(), "Used: %d"))

	leftCardFreeLabel := widget.NewLabel("")
	leftCardFreeLabel.Bind(binding.IntToStringWithFormat(model.LeftCardFree(), "Free: %d"))

	// Create labels for right card statistics
	rightCardTotalLabel := widget.NewLabel("")
	rightCardTotalLabel.Bind(binding.IntToStringWithFormat(model.RightCardTotal(), "Total: %d"))

	rightCardUsedLabel := widget.NewLabel("")
	rightCardUsedLabel.Bind(binding.IntToStringWithFormat(model.RightCardUsed(), "Used: %d"))

	rightCardFreeLabel := widget.NewLabel("")
	rightCardFreeLabel.Bind(binding.IntToStringWithFormat(model.RightCardFree(), "Free: %d"))

	// Create formatted text labels that combine the statistics
	leftCardText := widget.NewLabel("")
	rightCardText := widget.NewLabel("")

	// Update formatted text when bindings change
	updateLeftText := func() {
		total, _ := model.LeftCardTotal().Get()
		if total > 0 {
			used, _ := model.LeftCardUsed().Get()
			free, _ := model.LeftCardFree().Get()
			leftCardText.SetText(fmt.Sprintf("Card 1: Total: %d | Used: %d | Free: %d", total, used, free))
		} else {
			leftCardText.SetText("Card 1: No card loaded")
		}
	}

	updateRightText := func() {
		total, _ := model.RightCardTotal().Get()
		if total > 0 {
			used, _ := model.RightCardUsed().Get()
			free, _ := model.RightCardFree().Get()
			rightCardText.SetText(fmt.Sprintf("Card 2: Total: %d | Used: %d | Free: %d", total, used, free))
		} else {
			rightCardText.SetText("Card 2: No card loaded")
		}
	}

	// Listen to binding changes
	model.LeftCardTotal().AddListener(binding.NewDataListener(updateLeftText))
	model.LeftCardUsed().AddListener(binding.NewDataListener(updateLeftText))
	model.LeftCardFree().AddListener(binding.NewDataListener(updateLeftText))
	model.RightCardTotal().AddListener(binding.NewDataListener(updateRightText))
	model.RightCardUsed().AddListener(binding.NewDataListener(updateRightText))
	model.RightCardFree().AddListener(binding.NewDataListener(updateRightText))

	// Initial update
	updateLeftText()
	updateRightText()

	// Create footer container with horizontal layout
	footerContent := container.NewHBox(
		leftCardText,
		layout.NewSpacer(),
		rightCardText,
	)

	// Add visual styling: top border and background for contrast
	// Create a background rectangle with light gray tone for better readability
	// Use a light gray that provides contrast but is easy to read
	backgroundRect := canvas.NewRectangle(color.RGBA{
		R: 240,
		G: 240,
		B: 240,
		A: 255,
	})
	
	// Create a top border line
	borderColor := color.RGBA{R: 180, G: 180, B: 180, A: 255} // Light gray border
	topBorder := canvas.NewLine(borderColor)
	topBorder.StrokeWidth = 1
	
	// Create a container with padding for better spacing
	paddedContent := container.NewPadded(footerContent)
	
	// Create border container with top border line
	borderedContent := container.NewBorder(
		topBorder, // Top border line
		nil,
		nil,
		nil,
		paddedContent,
	)
	
	// Stack background and bordered content
	view.container = container.NewStack(
		backgroundRect,
		borderedContent,
	)

	return view
}

// Container returns the container for this view.
func (v *MemoryCardBlockStatsView) Container() *fyne.Container {
	return v.container
}

// UpdateStatistics triggers an update of the statistics from the memory cards.
func (v *MemoryCardBlockStatsView) UpdateStatistics() {
	v.model.UpdateStatistics()
}
