package blocks

import (
	"image/color"
	"reflect"

	"com.yvka.memcard/pkg/memcard"
	animatedsprite "com.yvka.memcard/pkg/ui/animated-sprite"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	SELECTED_BORDER_COLOR   = color.RGBA{R: 200, G: 100, B: 100, A: 255}
	UNSELECTED_BORDER_COLOR = color.RGBA{R: 0, G: 0, B: 00, A: 60}
	FILL_COLOR              = color.RGBA{R: 100, G: 100, B: 200, A: 255}
	BLOCK_SIZE              = float32(64.0)
)

type blockView struct {
	widget.BaseWidget
	model         BlockModelView
	container     *fyne.Container
	block         *canvas.Rectangle
	iconContainer *fyne.Container
}

func NewBlockView(idx int, cardId memcard.MemoryCardID, blockSelector BlockSelector) *blockView {

	model := NewBlockModelView(idx, cardId, blockSelector)

	block := canvas.NewRectangle(FILL_COLOR)
	block.StrokeColor = UNSELECTED_BORDER_COLOR
	block.StrokeWidth = 2
	block.SetMinSize(fyne.NewSize(BLOCK_SIZE, BLOCK_SIZE))
	block.Resize(fyne.NewSize(BLOCK_SIZE, BLOCK_SIZE))

	model.Allocated.AddListener(binding.NewDataListener(func() {
		allocated, _ := model.Allocated.Get()
		if allocated {
			block.FillColor = color.RGBA{
				R: FILL_COLOR.R,
				G: FILL_COLOR.G,
				B: FILL_COLOR.B,
				A: 255,
			}
		} else {
			block.FillColor = color.RGBA{
				R: FILL_COLOR.R,
				G: FILL_COLOR.G,
				B: FILL_COLOR.B,
				A: 0,
			}
		}
	}))

	blockLayout := container.NewStack(
		block,
	)

	bl := &blockView{
		model:     model,
		container: blockLayout,
		block:     block,
	}

	bl.ExtendBaseWidget(bl)

	bl.setupSelectedBinding()
	bl.setupIconAnimation()

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

func (v *blockView) setupIconAnimation() {
	model := v.model
	model.Animation.AddListener(binding.NewDataListener(func() {
		animation, _ := model.Animation.Get()

		// Remove existing icon if icon binding is nil
		if v.iconContainer != nil && reflect.ValueOf(animation).IsZero() {
			v.container.Remove(v.iconContainer)
			v.iconContainer = nil
		}

		if !reflect.ValueOf(animation).IsZero() {
			image := animatedsprite.NewAnimatedSprite(animation)
			if v.iconContainer != nil {
				v.container.Remove(v.iconContainer)
			}

			v.iconContainer = container.NewPadded(&image.Image)

			v.container.Add(v.iconContainer)
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

func (v *blockView) SetAnimation(animation animatedsprite.Animation) {
	v.model.Animation.Set(animation)
}

func (v *blockView) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(v.container)
}
