package blockstats

import (
	"fmt"

	"fyne.io/fyne/v2"
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
	footerContainer := container.NewHBox(
		leftCardText,
		layout.NewSpacer(),
		rightCardText,
	)

	view.container = footerContainer

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
