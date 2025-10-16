package memcard

import (
	"image"
	"image/color"
)

// Icon Image is 16x16 pixels, 4 bits per pixel (16 colors), so 128 bytes per frame
type IconBitmapFrame [128]byte

func (ib *IconBitmapFrame) PixelAt(x, y int) byte {
	if x < 0 || x >= 16 || y < 0 || y >= 16 {
		return 0
	}
	byteIndex := y*8 + x/2
	pixelData := ib[byteIndex]
	if x%2 == 0 {
		return pixelData & 0x0F // Lower nibble
	}
	return (pixelData >> 4) & 0x0F // Upper nibble
}

func (ib *IconBitmapFrame) ToImage(IconColorPalette [16]uint16) image.Image {

	paletteColors := color.Palette{}

	for _, c := range IconColorPalette {
		r := uint8((c & 0x1F) << 3)
		g := uint8(((c >> 5) & 0x1F) << 3)
		b := uint8(((c >> 10) & 0x1F) << 3)
		paletteColors = append(paletteColors, color.RGBA{R: r, G: g, B: b, A: 255})
	}

	img := image.NewPaletted(image.Rect(0, 0, 16, 16), paletteColors)

	for y := range 16 {
		for x := range 16 {
			colorIdx := ib.PixelAt(x, y)
			img.SetColorIndex(x, y, colorIdx)
		}
	}
	return img

}
