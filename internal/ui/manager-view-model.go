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
	} else {
		blockBindingList = vm.blocksRight
		vm.rightMemoryCard = card
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

func (vm *ManagerWindowViewModel) CopyCommand(sourceCardId memcard.MemoryCardID, blockIndex int) error {
	card := vm.getMemoryCardById(sourceCardId)

	if blockIndex < 0 || blockIndex >= memcard.NumBlocks {
		return fmt.Errorf("cannot copy block without selecting a block")
	}

	if card == nil {
		return fmt.Errorf("cannot copy block without loading a memory card \"%s\"", sourceCardId)
	}

	if err := card.CopyBlockTo(blockIndex, card); err != nil {
		return err
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

	// TODO: Have a method that refresh the bindings for all blocks based on the changed memory card
	var blockBindingList binding.UntypedList
	if sourceCardId == memcard.MemoryCardLeft {
		blockBindingList = vm.blocksLeft
		vm.leftMemoryCard = card
	} else {
		blockBindingList = vm.blocksRight
		vm.rightMemoryCard = card
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
	card := vm.getMemoryCardById(cardId)
	if card == nil {
		vm.setDefaultSaveGameTitle(cardId, blockIndex)
		return
	}

	blockItem, err := card.GetBlock(blockIndex)
	if err != nil || blockItem == nil {
		vm.setDefaultSaveGameTitle(cardId, blockIndex)
		return
	}

	vm.selectedSaveGameTitle.Set(blockItem.Title)

}

func (vm *ManagerWindowViewModel) setDefaultSaveGameTitle(cardId memcard.MemoryCardID, blockIndex int) {
	vm.selectedSaveGameTitle.Set(fmt.Sprintf("Card %s - Block %d", cardId, blockIndex))
}
