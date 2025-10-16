package ui

import (
	"fmt"

	"com.yv35.memcard/internal/memcard"
	_ui_blocks "com.yv35.memcard/internal/ui/blocks"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
)

type ManagerWindowViewModel struct {
	window fyne.Window

	selectedBlockIndex binding.Int
	selectedCardId     binding.String

	selectedSaveGameTitle binding.String

	blocksLeft  binding.UntypedList
	blocksRight binding.UntypedList
}

func NewManagerWindowViewModel(window fyne.Window) *ManagerWindowViewModel {
	return &ManagerWindowViewModel{
		window:                window,
		selectedBlockIndex:    binding.NewInt(),
		selectedCardId:        binding.NewString(),
		blocksLeft:            binding.NewUntypedList(),
		blocksRight:           binding.NewUntypedList(),
		selectedSaveGameTitle: binding.NewString(),
	}
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
	} else {
		blockBindingList = vm.blocksRight
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

func (vm *ManagerWindowViewModel) CopyCommand(source memcard.MemoryCardID, blockIndex int) {
}

func (vm *ManagerWindowViewModel) CopyAllCommand(source memcard.MemoryCardID) {
}

func (vm *ManagerWindowViewModel) DeleteCommand(source memcard.MemoryCardID, blockIndex int) {
}

func (vm *ManagerWindowViewModel) SelectedBlockIndex() int {
	val, err := vm.selectedBlockIndex.Get()
	if err != nil {
		return -1
	}
	return val
}

func (vm *ManagerWindowViewModel) SelectedCard() memcard.MemoryCardID {
	val, err := vm.selectedCardId.Get()
	if err != nil {
		return ""
	}
	return memcard.MemoryCardID(val)

}

func (vm *ManagerWindowViewModel) HandleBlockSelectionChanged(cardId memcard.MemoryCardID, blockIndex int) {
	vm.selectedCardId.Set(string(cardId))
	vm.selectedBlockIndex.Set(blockIndex)

	vm.selectedSaveGameTitle.Set(fmt.Sprintf("Selected: Card %s - Block %d", cardId, blockIndex))
}
