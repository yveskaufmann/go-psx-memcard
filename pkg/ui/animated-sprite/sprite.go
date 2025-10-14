package animatedsprite

import (
	"image"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Animation struct {
	Frames     []image.Image
	Loop       bool
	FrameDelay int64 // in milliseconds
}

func NewAnimation(frames []image.Image) Animation {

	frameDelay := int64(100)
	switch {
	case len(frames) == 2:
		frameDelay = 17 * 11 // 186ms per frame
	case len(frames) == 3:
		frameDelay = 17 * 16 // ~272ms per frame
	}

	return Animation{
		Frames:     frames,
		Loop:       len(frames) > 1,
		FrameDelay: frameDelay,
	}
}

type AnimatedSprite struct {
	Animation    Animation
	currentFrame int
	canvas.Image
}

func NewAnimatedSprite(animation Animation) *AnimatedSprite {
	sprite := &AnimatedSprite{}
	sprite.currentFrame = 0

	sprite.SetAnimation(animation)

	ticker := time.NewTicker(time.Millisecond * time.Duration(animation.FrameDelay))
	go func() {
		for range ticker.C {
			if sprite.currentFrame < len(sprite.Animation.Frames)-1 {
				sprite.currentFrame++
				fyne.Do(func() {
					sprite.Refresh()
				})
			} else if sprite.Animation.Loop {
				sprite.currentFrame = 0
				fyne.Do(func() {
					sprite.Refresh()
				})
			} else {
				ticker.Stop()
				return
			}
		}
	}()

	return sprite
}

func (s *AnimatedSprite) Refresh() {
	s.Image.Image = s.Animation.Frames[s.currentFrame]
	s.Image.FillMode = canvas.ImageFillContain
	s.Image.Refresh()
}

func (s *AnimatedSprite) SetAnimation(animation Animation) {
	s.Animation = animation
	s.currentFrame = 0
	s.Refresh()
}
