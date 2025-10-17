package blocks

import (
	"com.yv35.memcard/internal/memcard"
	animatedsprite "com.yv35.memcard/internal/ui/animated-sprite"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type BlockSelector interface {
	SelectBlock(idx int)
	UnselectBlock(idx int)
}
type BlockContainer struct {
	widget.BaseWidget
	blocks       []*blockView
	blockBinding binding.UntypedList
	// selectedBlockIndexes removed - selection state is now managed by the ViewModel
	onBlockSelected func(blockIndex int)
}

func NewBlockContainer(cardId memcard.MemoryCardID, blockBinding binding.UntypedList) *BlockContainer {
	bc := &BlockContainer{
		blockBinding: blockBinding,
	}
	bc.ExtendBaseWidget(bc)

	for i := 0; i < 15; i++ {
		block := NewBlockView(i, cardId, bc)
		bc.blocks = append(bc.blocks, block)
	}

	blockBinding.AddListener(binding.NewDataListener(func() {
		bc.Refresh()
	}))

	return bc
}

func (b *BlockContainer) CreateRenderer() fyne.WidgetRenderer {
	grid := container.NewGridWithColumns(3)

	for _, block := range b.blocks {
		grid.Add(block)
	}

	return widget.NewSimpleRenderer(grid)
}

func (b *BlockContainer) SetOnBlockSelected(callback func(blockIndex int)) {
	b.onBlockSelected = callback
}

func (b *BlockContainer) Refresh() {
	b.BaseWidget.Refresh()

	// Update the block views based on the current state of the blocks list

	for i := range len(b.blocks) {
		b.blocks[i].Unselect()
		b.blocks[i].model.Allocated.Set(false)
		b.blocks[i].model.Animation.Set(animatedsprite.Animation{})
		b.blocks[i].model.GameTitle.Set("")
	}

	for i := 0; i < b.blockBinding.Length(); i++ {
		item, err := b.blockBinding.GetItem(i)
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

			b.blocks[idx].SetAnimation(animation)
			b.blocks[idx].model.Allocated.Set(true)

		}
	}
}

func (b *BlockContainer) SelectBlock(idx int) {
	if idx < 0 || idx >= len(b.blocks) {
		return
	}

	block := b.blocks[idx]
	if block.Selected() {
		return
	}
	
	if b.onBlockSelected != nil {
		b.onBlockSelected(idx)
	}

}

func (b *BlockContainer) UnselectBlock(idx int) {
	if idx < 0 || idx >= len(b.blocks) {
		return
	}

	block := b.blocks[idx]
	if !block.Selected() {
		return
	}
	block.Unselect()

	if b.onBlockSelected != nil {
		b.onBlockSelected(-1)
	}
}

// ClearAllSelections clears the visual selection state of all blocks in this container
func (b *BlockContainer) ClearAllSelections() {
	for _, block := range b.blocks {
		if block.Selected() {
			block.Unselect()
		}
	}
}

func (b *BlockContainer) SetBlockItem(idx int, item memcard.BlockItem) {
	if idx < 0 || idx >= len(b.blocks) {
		return
	}

	block := b.blocks[idx]
	block.SetAnimation(item.Animation)
}

type Item struct {
	Index     int
	Title     string
	Animation animatedsprite.Animation
	Used      bool
}
