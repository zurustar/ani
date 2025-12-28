package renderer

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/fogleman/gg"
)

// RenderFrame draws a single frame with the image at (x, y).
// It returns a paletted image with transparency, using the provided palette.
func RenderFrame(img image.Image, width, height, x, y int, p color.Palette) *image.Paletted {
	dc := gg.NewContext(width, height)

	// Create a transparent background
	dc.SetColor(color.Transparent)
	dc.Clear()

	// Draw the image
	dc.DrawImage(img, x, y)

	bounds := image.Rect(0, 0, width, height)
	dst := image.NewPaletted(bounds, p)

	// Draw using draw.Src to preserve exact colors where possible
	draw.Draw(dst, bounds, dc.Image(), image.Point{}, draw.Src)

	return dst
}
