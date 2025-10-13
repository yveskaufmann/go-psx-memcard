package ui

import (
	"fmt"

	"com.yvka.memcard/pkg/memcard"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
)

type ManagerWindowViewModel struct {
	window fyne.Window

	selectedBlockIndex binding.Int
	selectedCardId     binding.String

	blocksLeft  binding.UntypedList
	blocksRight binding.UntypedList
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

	vm.blocksLeft = binding.NewUntypedList()
	vm.blocksRight = binding.NewUntypedList()

	for idx, block := range blocks {
		leftMemoryCardView.SetBlockItem(idx, block)

	}

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
}
