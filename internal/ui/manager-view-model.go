package ui

import (
	"fmt"

	"com.yv35.memcard/internal/memcard"
	_ui_blocks "com.yv35.memcard/internal/ui/blocks"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
)

const NoBlockSelected = -1

type ManagerWindowViewModel struct {
	window fyne.Window

	selection *_ui_blocks.SelectionViewModel

	selectedSaveGameTitle binding.String

	blocksLeft  binding.UntypedList
	blocksRight binding.UntypedList

	leftMemoryCard  *memcard.MemoryCard
	rightMemoryCard *memcard.MemoryCard

	leftMemoryCardPath  string
	rightMemoryCardPath string
}

func NewManagerWindowViewModel(window fyne.Window) *ManagerWindowViewModel {
	win := &ManagerWindowViewModel{
		window:                window,
		blocksLeft:            binding.NewUntypedList(),
		blocksRight:           binding.NewUntypedList(),
		selectedSaveGameTitle: binding.NewString(),
		selection:             _ui_blocks.NewBlockSelectionViewModel(),
	}

	win.selectedSaveGameTitle.Set("")

	return win
}

func (vm *ManagerWindowViewModel) SelectionViewModel() *_ui_blocks.SelectionViewModel {
	return vm.selection
}

func (vm *ManagerWindowViewModel) LoadMemoryCardImage(path string, memoryCardId memcard.MemoryCardID) {
	// Open the memory card file
	card, err := memcard.Open(path)
	if err != nil {
		dialog.ShowError(err, vm.window)
		return
	}

	fmt.Printf("Loaded memory card: %+v\n", card)
	blocks, err := card.ListBlocks()
	if err != nil {
		dialog.ShowError(err, vm.window)
		return
	}

	var blockBindingList binding.UntypedList
	if memoryCardId == memcard.MemoryCardLeft {
		blockBindingList = vm.blocksLeft
		vm.leftMemoryCard = card
		vm.leftMemoryCardPath = path
	} else {
		blockBindingList = vm.blocksRight
		vm.rightMemoryCard = card
		vm.rightMemoryCardPath = path
	}

	bindings := []any{}
	for idx, block := range blocks {

		blockItem := _ui_blocks.Item{
			Index:     idx,
			Title:     block.Title,
			Animation: block.Animation,
			Used:      block.Title != "",
		}

		bindings = append(bindings, blockItem)
	}

	blockBindingList.Set(bindings)
}

func (vm *ManagerWindowViewModel) getMemoryCardById(cardId memcard.MemoryCardID) *memcard.MemoryCard {
	if cardId == memcard.MemoryCardLeft {
		return vm.leftMemoryCard
	}
	return vm.rightMemoryCard
}

func (vm *ManagerWindowViewModel) GetOppositeMemoryCardId(cardId memcard.MemoryCardID) memcard.MemoryCardID {
	if cardId == memcard.MemoryCardLeft {
		return memcard.MemoryCardRight
	}
	return memcard.MemoryCardLeft
}

func (vm *ManagerWindowViewModel) GetMemoryCardPathById(cardId memcard.MemoryCardID) string {
	if cardId == memcard.MemoryCardLeft {
		return vm.leftMemoryCardPath
	}
	return vm.rightMemoryCardPath
}

func (vm *ManagerWindowViewModel) CopyCommand(sourceCardId memcard.MemoryCardID, blockIndex int) error {
	sourceCard := vm.getMemoryCardById(sourceCardId)

	if blockIndex < 0 || blockIndex >= memcard.NumBlocks {
		return fmt.Errorf("cannot copy block without selecting a block")
	}

	if sourceCard == nil {
		return fmt.Errorf("cannot copy block without loading a memory card \"%s\"", sourceCardId)
	}

	// Get the target card (opposite card)
	targetCardId := vm.GetOppositeMemoryCardId(sourceCardId)
	targetCard := vm.getMemoryCardById(targetCardId)

	if targetCard == nil {
		return fmt.Errorf("cannot copy block: target memory card \"%s\" is not loaded", targetCardId)
	}

	// Copy the block to the target card
	if err := sourceCard.CopyBlockTo(blockIndex, targetCard); err != nil {
		return fmt.Errorf("failed to copy block: %w", err)
	}

	// Write the target card to disk
	targetCardPath := vm.GetMemoryCardPathById(targetCardId)
	if err := targetCard.Write(targetCardPath); err != nil {
		return fmt.Errorf("failed to write target memory card: %w", err)
	}

	// Refresh the target card bindings to show the new block
	if err := vm.RefreshCardBindings(targetCardId); err != nil {
		return fmt.Errorf("failed to refresh target card bindings: %w", err)
	}

	return nil
}

func (vm *ManagerWindowViewModel) DeleteCommand(sourceCardId memcard.MemoryCardID, blockIndex int) error {
	card := vm.getMemoryCardById(sourceCardId)

	if blockIndex < 0 || blockIndex >= memcard.NumBlocks {
		return fmt.Errorf("cannot delete block without selecting a block")
	}

	if card == nil {
		return fmt.Errorf("cannot delete block without loading a memory card \"%s\"", sourceCardId)
	}

	if err := card.DeleteBlockFrom(blockIndex); err != nil {
		return err
	}

	vm.RefreshCardBindings(sourceCardId)

	return card.Write(vm.GetMemoryCardPathById(sourceCardId))
}

func (vm *ManagerWindowViewModel) RefreshCardBindings(sourceCardId memcard.MemoryCardID) error {

	card := vm.getMemoryCardById(sourceCardId)
	if card == nil {
		return fmt.Errorf("cannot refresh bindings without loading a memory card \"%s\"", sourceCardId)
	}

	// TODO: Have a method that refresh the bindings for all blocks based on the changed memory card
	var blockBindingList binding.UntypedList
	if sourceCardId == memcard.MemoryCardLeft {
		blockBindingList = vm.blocksLeft
	} else {
		blockBindingList = vm.blocksRight
	}

	blocks, err := card.ListBlocks()
	if err != nil {
		dialog.ShowError(err, vm.window)
		return err
	}

	bindings := []any{}
	for idx, block := range blocks {

		blockItem := _ui_blocks.Item{
			Index:     idx,
			Title:     block.Title,
			Animation: block.Animation,
			Used:      block.Title != "",
		}

		bindings = append(bindings, blockItem)
	}

	blockBindingList.Set(bindings)
	return nil
}

func (vm *ManagerWindowViewModel) SelectedBlockIndex() int {
	return vm.selection.BlockIndex()
}

func (vm *ManagerWindowViewModel) SelectedCard() memcard.MemoryCardID {
	return vm.selection.CardId()
}

func (vm *ManagerWindowViewModel) HandleBlockSelectionChanged(cardId memcard.MemoryCardID, blockIndex int) {
	// Clear title if no block is selected
	if blockIndex == NoBlockSelected || cardId == "" {
		vm.selectedSaveGameTitle.Set("")
		return
	}

	card := vm.getMemoryCardById(cardId)
	if card == nil {
		vm.selectedSaveGameTitle.Set("")
		return
	}

	blockItem, err := card.GetBlock(blockIndex)
	if err != nil || blockItem == nil {
		vm.selectedSaveGameTitle.Set("")
		return
	}

	vm.selectedSaveGameTitle.Set(blockItem.Title)

}

func (vm *ManagerWindowViewModel) setDefaultSaveGameTitle(cardId memcard.MemoryCardID, blockIndex int) {
	vm.selectedSaveGameTitle.Set(fmt.Sprintf("Card %s - Block %d", cardId, blockIndex))
}

// GetBlockStatistics returns the total, used, and free block counts for the specified memory card.
// Returns (0, 0, 0) if the card is not loaded.
func (vm *ManagerWindowViewModel) GetBlockStatistics(cardId memcard.MemoryCardID) (total, used, free int) {
	card := vm.getMemoryCardById(cardId)
	if card == nil {
		return 0, 0, 0
	}
	return card.CountBlocks()
}
