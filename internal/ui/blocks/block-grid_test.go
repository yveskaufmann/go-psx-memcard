package blocks

import (
	"testing"

	"com.yv35.memcard/internal/memcard"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"
)

// TestBlockContainer_SelectBlock_ClearsSiblingSelection tests that selecting a block
// in one container clears the selection in its sibling container
func TestBlockContainer_SelectBlock_ClearsSiblingSelection(t *testing.T) {
	// Initialize test app
	_ = test.NewTempApp(t)

	// Create two block containers (simulating left and right memory cards)
	leftBinding := binding.NewUntypedList()
	rightBinding := binding.NewUntypedList()

	leftContainer := NewBlockContainer(memcard.MemoryCardLeft, leftBinding)
	rightContainer := NewBlockContainer(memcard.MemoryCardRight, rightBinding)

	// Set up sibling relationship
	leftContainer.SetSiblingContainer(rightContainer)
	rightContainer.SetSiblingContainer(leftContainer)

	// Select a block in the left container
	leftContainer.SelectBlock(0)

	// Verify the block is selected in left container
	if len(leftContainer.selectedBlockIndexes) != 1 || leftContainer.selectedBlockIndexes[0] != 0 {
		t.Errorf("Expected block 0 to be selected in left container, got %v", leftContainer.selectedBlockIndexes)
	}

	// Verify no blocks are selected in right container
	if len(rightContainer.selectedBlockIndexes) != 0 {
		t.Errorf("Expected no blocks to be selected in right container, got %v", rightContainer.selectedBlockIndexes)
	}

	// Select a block in the right container
	rightContainer.SelectBlock(5)

	// Verify the block is selected in right container
	if len(rightContainer.selectedBlockIndexes) != 1 || rightContainer.selectedBlockIndexes[0] != 5 {
		t.Errorf("Expected block 5 to be selected in right container, got %v", rightContainer.selectedBlockIndexes)
	}

	// Verify the left container's selection was cleared
	if len(leftContainer.selectedBlockIndexes) != 0 {
		t.Errorf("Expected left container selection to be cleared, got %v", leftContainer.selectedBlockIndexes)
	}
}

// TestBlockContainer_ClearSelection tests that ClearSelection properly clears all selections
func TestBlockContainer_ClearSelection(t *testing.T) {
	// Initialize test app
	_ = test.NewTempApp(t)

	binding := binding.NewUntypedList()
	container := NewBlockContainer(memcard.MemoryCardLeft, binding)

	// Select a block
	container.SelectBlock(2)

	// Verify block is selected
	if len(container.selectedBlockIndexes) != 1 {
		t.Errorf("Expected 1 selected block, got %d", len(container.selectedBlockIndexes))
	}

	// Clear selection
	container.ClearSelection()

	// Verify selection is cleared
	if len(container.selectedBlockIndexes) != 0 {
		t.Errorf("Expected no selected blocks after ClearSelection, got %v", container.selectedBlockIndexes)
	}
}

// TestBlockContainer_SelectBlock_WithinSameContainer tests that selecting a different block
// within the same container clears the previous selection
func TestBlockContainer_SelectBlock_WithinSameContainer(t *testing.T) {
	// Initialize test app
	_ = test.NewTempApp(t)

	binding := binding.NewUntypedList()
	container := NewBlockContainer(memcard.MemoryCardLeft, binding)

	// Select first block
	container.SelectBlock(0)

	// Verify first block is selected
	if len(container.selectedBlockIndexes) != 1 || container.selectedBlockIndexes[0] != 0 {
		t.Errorf("Expected block 0 to be selected, got %v", container.selectedBlockIndexes)
	}

	// Select second block
	container.SelectBlock(3)

	// Verify only second block is selected
	if len(container.selectedBlockIndexes) != 1 || container.selectedBlockIndexes[0] != 3 {
		t.Errorf("Expected only block 3 to be selected, got %v", container.selectedBlockIndexes)
	}
}
