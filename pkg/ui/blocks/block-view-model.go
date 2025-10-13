package blocks

import (
	"image"

	"com.yvka.memcard/pkg/memcard"
	"fyne.io/fyne/v2/data/binding"
)

type BlockModelView struct {
	Index         int
	CardId        memcard.MemoryCardID
	Selected      binding.Bool
	Allocated     binding.Bool
	GameTitle     binding.String            // binding to string
	Icon          binding.Item[image.Image] // binding to image.Image
	blockSelector BlockSelector
}

func NewBlockModelView(idx int, cardId memcard.MemoryCardID, blockSelector BlockSelector) BlockModelView {
	return BlockModelView{
		Index:         idx,
		CardId:        cardId,
		Selected:      binding.NewBool(),
		blockSelector: blockSelector,
		Allocated:     binding.NewBool(),
		GameTitle:     nil,
		Icon:          nil,
	}
}

func (b *BlockModelView) IsSelected() bool {
	val, err := b.Selected.Get()
	if err != nil {
		return false
	}
	return val
}

func (b *BlockModelView) ToggleSelect() {

	selected := b.IsSelected()

	if selected {
		b.blockSelector.UnselectBlock(b.Index)
	} else {
		b.blockSelector.SelectBlock(b.Index)
	}

	b.Selected.Set(!selected)
}

func (b *BlockModelView) UnSelect() {
	if !b.IsSelected() {
		return
	}
	b.Selected.Set(false)
	b.blockSelector.UnselectBlock(b.Index)
}
