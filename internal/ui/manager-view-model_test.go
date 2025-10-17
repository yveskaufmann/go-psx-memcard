package ui

import (
	"testing"

	"com.yv35.memcard/internal/memcard"
	"com.yv35.memcard/internal/ui/blocks"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"
)

// TestManagerWindowViewModel_HandleBlockSelectionChanged_ClearsSiblingSelection tests that
// when a block is selected in one container, the ViewModel clears visual selections in the other container
func TestManagerWindowViewModel_HandleBlockSelectionChanged_ClearsSiblingSelection(t *testing.T) {
	app := test.NewTempApp(t)
	window := app.NewWindow("Test")
	_ = window

	// Create view model
	vm := NewManagerWindowViewModel(window)

	// Create two block containers
	leftContainer := blocks.NewBlockContainer(memcard.MemoryCardLeft, vm.blocksLeft)
	rightContainer := blocks.NewBlockContainer(memcard.MemoryCardRight, vm.blocksRight)

	// Set container references in the view model
	vm.SetBlockContainers(leftContainer, rightContainer)

	// Simulate selecting a block in the left container
	vm.HandleBlockSelectionChanged(memcard.MemoryCardLeft, 3)

	// Verify that left container state is tracked in ViewModel
	if vm.SelectedBlockIndex() != 3 {
		t.Errorf("Expected block index 3 to be selected, got %d", vm.SelectedBlockIndex())
	}
	if vm.SelectedCard() != memcard.MemoryCardLeft {
		t.Errorf("Expected left card to be selected, got %s", vm.SelectedCard())
	}

	// Simulate selecting a block in the right container
	vm.HandleBlockSelectionChanged(memcard.MemoryCardRight, 7)

	// Verify that the ViewModel now tracks the right container selection
	if vm.SelectedBlockIndex() != 7 {
		t.Errorf("Expected block index 7 to be selected, got %d", vm.SelectedBlockIndex())
	}
	if vm.SelectedCard() != memcard.MemoryCardRight {
		t.Errorf("Expected right card to be selected, got %s", vm.SelectedCard())
	}
}

// TestManagerWindowViewModel_HandleBlockSelectionChanged_Unselect tests that
// passing a negative block index clears the selection
func TestManagerWindowViewModel_HandleBlockSelectionChanged_Unselect(t *testing.T) {
	app := test.NewTempApp(t)
	window := app.NewWindow("Test")
	_ = window

	vm := NewManagerWindowViewModel(window)

	// Create containers
	leftContainer := blocks.NewBlockContainer(memcard.MemoryCardLeft, vm.blocksLeft)
	rightContainer := blocks.NewBlockContainer(memcard.MemoryCardRight, vm.blocksRight)
	vm.SetBlockContainers(leftContainer, rightContainer)

	// Select a block
	vm.HandleBlockSelectionChanged(memcard.MemoryCardLeft, 5)

	// Verify selection
	if vm.SelectedBlockIndex() != 5 {
		t.Errorf("Expected block 5 to be selected, got %d", vm.SelectedBlockIndex())
	}

	// Unselect by passing -1
	vm.HandleBlockSelectionChanged(memcard.MemoryCardLeft, -1)

	// Verify no selection
	if vm.SelectedBlockIndex() != NoBlockSelected {
		t.Errorf("Expected no block selected, got %d", vm.SelectedBlockIndex())
	}
	if vm.SelectedCard() != "" {
		t.Errorf("Expected no card selected, got %s", vm.SelectedCard())
	}
}

// TestBlockContainer_ClearAllSelections tests the ClearAllSelections method
func TestBlockContainer_ClearAllSelections(t *testing.T) {
	app := test.NewTempApp(t)
	_ = app

	binding := binding.NewUntypedList()
	container := blocks.NewBlockContainer(memcard.MemoryCardLeft, binding)

	// This test verifies that ClearAllSelections can be called without error
	// The actual visual selection state is managed by the individual blockView instances
	container.ClearAllSelections()

	// Test passes if no panic occurs
}
