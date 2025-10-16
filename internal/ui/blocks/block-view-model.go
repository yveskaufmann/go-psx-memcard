package blocks

import (
	"com.yv35.memcard/internal/memcard"
	animatedsprite "com.yv35.memcard/internal/ui/animated-sprite"
	"fyne.io/fyne/v2/data/binding"
)

type BlockModelView struct {
	Index         int
	CardId        memcard.MemoryCardID
	Selected      binding.Bool
	Allocated     binding.Bool
	GameTitle     binding.String                         // binding to string
	Animation     binding.Item[animatedsprite.Animation] // binding to animatedsprite.Animation
	blockSelector BlockSelector
}

func NewBlockModelView(idx int, cardId memcard.MemoryCardID, blockSelector BlockSelector) BlockModelView {
	return BlockModelView{
		Index:         idx,
		CardId:        cardId,
		Selected:      binding.NewBool(),
		blockSelector: blockSelector,
		Allocated:     binding.NewBool(),
		GameTitle:     binding.NewString(),
		Animation:     binding.NewItem((func(a, b animatedsprite.Animation) bool { return len(a.Frames) == len(b.Frames) })),
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
