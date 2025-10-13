package blocks

import (
	"image"
	"image/color"

	"com.yvka.memcard/pkg/memcard"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	SELECTED_BORDER_COLOR   = color.RGBA{R: 200, G: 100, B: 100, A: 255}
	UNSELECTED_BORDER_COLOR = color.Black
	BLOCK_SIZE              = float32(64.0)
)

type blockView struct {
	widget.BaseWidget
	model     BlockModelView
	container *fyne.Container
	block     *canvas.Rectangle
	icon      *canvas.Image
}

func NewBlockView(idx int, cardId memcard.MemoryCardID, blockSelector BlockSelector) *blockView {

	model := NewBlockModelView(idx, cardId, blockSelector)

	block := canvas.NewRectangle(color.RGBA{R: 100, G: 100, B: 200, A: 255})
	block.StrokeColor = UNSELECTED_BORDER_COLOR
	block.StrokeWidth = 2
	block.SetMinSize(fyne.NewSize(BLOCK_SIZE, BLOCK_SIZE))
	block.Resize(fyne.NewSize(BLOCK_SIZE, BLOCK_SIZE))

	blockLayout := container.NewStack(
		block,
	)

	bl := &blockView{
		model:     model,
		container: blockLayout,
		block:     block,
	}

	bl.setupSelectedBinding()
	bl.setupIconBinding()

	bl.ExtendBaseWidget(bl)

	return bl
}

func (v *blockView) setupSelectedBinding() {
	model := v.model
	model.Selected.AddListener(binding.NewDataListener(func() {
		selected := model.IsSelected()
		if selected {
			v.block.StrokeColor = SELECTED_BORDER_COLOR
		} else {
			v.block.StrokeColor = UNSELECTED_BORDER_COLOR
		}
		v.block.Refresh()
	}))
}

func (v *blockView) setupIconBinding() {
	model := v.model
	model.Icon.AddListener(binding.NewDataListener(func() {
		icon, _ := model.Icon.Get()

		// Remove existing icon if icon binding is nil
		if v.icon != nil && icon == nil {
			v.container.Remove(v.icon)
			v.icon = nil
		}

		if v.icon != nil {
			v.icon = canvas.NewImageFromImage(icon)
			v.container.Objects = append(v.container.Objects, v.icon)
		}

		if icon != nil {
			v.SetIcon(icon)
		}

		v.block.Refresh()
	}))
}

func (v *blockView) Selected() bool {
	return v.model.IsSelected()
}

func (v *blockView) Select() {
	v.model.ToggleSelect()
}

func (v *blockView) Unselect() {
	v.model.UnSelect()
}

func (v *blockView) Tapped(ev *fyne.PointEvent) {
	v.model.ToggleSelect()
}

func (v *blockView) SetIcon(image image.Image) {
	v.model.Icon.Set(image)
}

func (v *blockView) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(v.container)
}
