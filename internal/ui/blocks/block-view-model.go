package blocks

import (
	"com.yv35.memcard/internal/memcard"
	animatedsprite "com.yv35.memcard/internal/ui/animated-sprite"
	"fyne.io/fyne/v2/data/binding"
)

type BlockModelView struct {
	Index          int
	CardId         memcard.MemoryCardID
	Selected       binding.Bool
	Allocated      binding.Bool
	GameTitle      binding.String                         // binding to string
	Animation      binding.Item[animatedsprite.Animation] // binding to animatedsprite.Animation
	blockSelection *SelectionViewModel
}

func NewBlockModelView(idx int, cardId memcard.MemoryCardID, blockSelector *SelectionViewModel) *BlockModelView {

	model := &BlockModelView{
		Index:          idx,
		CardId:         cardId,
		Selected:       binding.NewBool(),
		blockSelection: blockSelector,
		Allocated:      binding.NewBool(),
		GameTitle:      binding.NewString(),
		Animation:      binding.NewItem((func(a, b animatedsprite.Animation) bool { return len(a.Frames) == len(b.Frames) })),
	}

	blockSelector.AddListener(NewSelectionChangedListener(model.handleSelectionChanged))

	return model
}

func (b *BlockModelView) handleSelectionChanged(cardID memcard.MemoryCardID, blockIndex int) {
	isCurrentBlockSelected := (cardID == b.CardId) && (blockIndex == b.Index)
	b.Selected.Set(isCurrentBlockSelected)
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
		b.blockSelection.UnselectBlock(b.CardId, b.Index)
	} else {
		b.blockSelection.SelectBlock(b.CardId, b.Index)
	}

}

func (b *BlockModelView) UnSelect() {
	b.blockSelection.ClearSelection()
}
