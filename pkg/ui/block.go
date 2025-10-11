package ui

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	SELECTED_BORDER_COLOR   = color.RGBA{R: 200, G: 100, B: 100, A: 255}
	UNSELECTED_BORDER_COLOR = color.Black
	BLOCK_SIZE              = float32(64.0)
)

type blockElement struct {
	widget.BaseWidget
	idx           int
	container     *fyne.Container
	rect          *canvas.Rectangle
	selected      bool
	cardId        int
	blockSelector BlockSelector
}

func NewBlock(idx int, cardId int, blockSelector BlockSelector) *blockElement {

	block := canvas.NewRectangle(color.RGBA{R: 100, G: 100, B: 200, A: 255})
	block.StrokeColor = UNSELECTED_BORDER_COLOR
	block.StrokeWidth = 2
	block.SetMinSize(fyne.NewSize(BLOCK_SIZE, BLOCK_SIZE))
	block.Resize(fyne.NewSize(BLOCK_SIZE, BLOCK_SIZE))

	blockLayout := container.NewStack(
		block,
	)

	bl := &blockElement{
		idx:           idx,
		container:     blockLayout,
		rect:          block,
		selected:      false,
		cardId:        cardId,
		blockSelector: blockSelector,
	}

	bl.ExtendBaseWidget(bl)

	return bl
}

func (b *blockElement) Selected() bool {
	return b.selected
}

func (b *blockElement) Select() {

	b.selected = !b.selected
	if b.selected {
		b.rect.StrokeColor = SELECTED_BORDER_COLOR
	} else {
		b.rect.StrokeColor = UNSELECTED_BORDER_COLOR
	}
	b.rect.Refresh()
}

func (b *blockElement) SetIcon(image image.Image) {
	img := canvas.NewImageFromImage(image)
	b.container.Objects = append(b.container.Objects, img)
}

func (b *blockElement) Unselect() {
	if !b.selected {
		return
	}
	// Select act as a toggle, so we can just call it
	// to unselect if already selected
	b.Select()
}

func (b *blockElement) Tapped(ev *fyne.PointEvent) {
	b.blockSelector.SelectBlock(b.idx)
}

func (b *blockElement) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(b.container)
}
