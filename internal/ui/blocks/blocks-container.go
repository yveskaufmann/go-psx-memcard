package blocks

import (
	"com.yv35.memcard/internal/memcard"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type Container struct {
	widget.BaseWidget
	vm        *ContainerViewModel
	blocks    []*blockView
	selection *SelectionViewModel
}

func NewContainer(cardId memcard.MemoryCardID, blockBinding binding.UntypedList, blockSelection *SelectionViewModel) *Container {
	bc := &Container{
		vm:        NewBlockGridContainerViewModel(cardId, blockBinding, blockSelection),
		selection: blockSelection,
	}

	bc.ExtendBaseWidget(bc)

	for i := 0; i < 15; i++ {
		block := NewBlockView(i, cardId, bc.vm.Blocks[i])
		bc.blocks = append(bc.blocks, block)
	}

	blockBinding.AddListener(binding.NewDataListener(bc.Refresh))

	return bc
}

func (b *Container) CreateRenderer() fyne.WidgetRenderer {
	grid := container.NewGridWithColumns(3)

	for _, block := range b.blocks {
		grid.Add(block)
	}

	return widget.NewSimpleRenderer(grid)
}

func (b *Container) SetOnBlockSelected(callback func(cardID memcard.MemoryCardID, blockIndex int)) {
	b.vm.OnBlockSelected = callback
}

func (b *Container) Refresh() {
	b.BaseWidget.Refresh()
	b.vm.Refresh()
}
