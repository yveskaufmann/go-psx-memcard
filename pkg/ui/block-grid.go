package ui

import (
	"com.yvka.memcard/pkg/memcard"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type BlockSelector interface {
	SelectBlock(idx int)
	UnselectBlock(idx int)
}
type BlockContainer struct {
	widget.BaseWidget
	blocks               []*blockElement
	selectedBlockIndexes []int
}

func NewBlockContainer(cardId int) *BlockContainer {
	bc := &BlockContainer{}
	bc.ExtendBaseWidget(bc)

	for i := 0; i < 15; i++ {
		block := NewBlock(i, cardId, bc)
		bc.blocks = append(bc.blocks, block)
	}

	return bc
}

func (b *BlockContainer) CreateRenderer() fyne.WidgetRenderer {
	grid := container.NewGridWithColumns(3)

	for _, block := range b.blocks {
		grid.Add(block)
	}

	return widget.NewSimpleRenderer(grid)
}

func (b *BlockContainer) SelectBlock(idx int) {
	if idx < 0 || idx >= len(b.blocks) {
		return
	}

	if len(b.selectedBlockIndexes) > 0 {
		for _, selectedBlockIdx := range b.selectedBlockIndexes {
			block := b.blocks[selectedBlockIdx]
			block.Unselect()
		}
	}

	block := b.blocks[idx]
	if block.Selected() {
		return
	}
	block.Select()
	b.selectedBlockIndexes = []int{idx}

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

	b.selectedBlockIndexes = sliceFilter(b.selectedBlockIndexes, func(i int) bool {
		return i != idx
	})

}

func (b *BlockContainer) SetBlockItem(idx int, item memcard.BlockItem) {
	if idx < 0 || idx >= len(b.blocks) {
		return
	}

	block := b.blocks[idx]
	block.SetIcon(item.Icon)
}

func sliceFilter(s []int, test func(int) bool) (ret []int) {
	for _, v := range s {
		if test(v) {
			ret = append(ret, v)
		}
	}
	return
}
