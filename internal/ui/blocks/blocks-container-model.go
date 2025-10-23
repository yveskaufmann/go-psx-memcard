package blocks

import (
	"com.yv35.memcard/internal/memcard"
	animatedsprite "com.yv35.memcard/internal/ui/animated-sprite"
	"fyne.io/fyne/v2/data/binding"
)

const TotalBlocksPerCard = 15

type ContainerViewModel struct {
	BlockBindings   binding.UntypedList
	OnBlockSelected func(cardId memcard.MemoryCardID, blockIndex int)
	cardId          memcard.MemoryCardID
	Blocks          []*BlockModelView
}

func NewBlockGridContainerViewModel(cardId memcard.MemoryCardID, blockBindings binding.UntypedList, blockSelection *SelectionViewModel) *ContainerViewModel {
	vm := &ContainerViewModel{
		cardId:        cardId,
		BlockBindings: blockBindings,
	}

	for i := range TotalBlocksPerCard {
		blockVM := NewBlockModelView(i, cardId, blockSelection)
		vm.Blocks = append(vm.Blocks, blockVM)
	}

	vm.BlockBindings.AddListener(binding.NewDataListener(func() {
		vm.Refresh()
	}))

	blockSelection.AddListener(NewSelectionChangedListener(func(cardID memcard.MemoryCardID, blockIndex int) {
		if vm.OnBlockSelected != nil {
			vm.OnBlockSelected(cardID, blockIndex)
		}
	}))

	return vm
}

func (c *ContainerViewModel) Refresh() {

	// Update the block views based on the current state of the blocks list
	for i := range len(c.Blocks) {
		c.Blocks[i].Allocated.Set(false)
		c.Blocks[i].Animation.Set(animatedsprite.Animation{})
		c.Blocks[i].GameTitle.Set("")
	}

	for i := 0; i < c.BlockBindings.Length(); i++ {
		item, err := c.BlockBindings.GetItem(i)
		if err != nil {
			continue
		}

		blockItem, ok := item.(binding.Untyped)
		if !ok {
			continue
		}

		v, err := blockItem.Get()
		if err != nil {
			continue
		}

		block, ok := v.(Item)
		if !ok {
			continue
		}

		if blockItem != nil {
			idx := block.Index
			animation := block.Animation

			c.Blocks[idx].Animation.Set(animation)
			c.Blocks[idx].Allocated.Set(true)

		}
	}
}

type Item struct {
	Index     int
	Title     string
	Animation animatedsprite.Animation
	Used      bool
}
