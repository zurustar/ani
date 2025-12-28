package renderer

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"

	"github.com/fogleman/gg"
)

// RenderFrame draws a single frame with the image at (x, y).
// It returns a paletted image with transparency.
func RenderFrame(img image.Image, width, height, x, y int) *image.Paletted {
	dc := gg.NewContext(width, height)

	// Create a transparent background
	dc.SetColor(color.Transparent)
	dc.Clear()

	// Draw the image
	dc.DrawImage(img, x, y)

	// Convert to Paletted image
	// We use Plan9 palette. We need to ensure one index is transparent.
	// Plan9 palette index 0 is NOT guaranteed perfectly transparent in all contexts if not set explicitly,
	// but `image/gif` handles pure transparency if it exists in the image.
	// However, `gg` writes to RGBA. When converting RGBA to Paletted, we need a palette that includes transparent.

	// Create a palette that includes transparent color if Plan9 doesn't handle it well for direct transparency mapping.
	// Actually, palette.Plan9 does not contain a fully transparent color (alpha 0) except generally implicitly handled.
	// A better approach is to use a custom palette or standard WebSafe, but let's try standard conversion first.
	// To ensure transparency works, we typically append a transparent color to the palette or pick one.

	// Simplified approach for V1:
	// 1. Get RGBA image from context.
	// 2. Conver to Paletted using a palette that has transparency.

	// Let's use palette.Plan9 and modify index 0 to be transparent?
	// Or simpler: Quantize? No, design said standard palette.

	// We use palette.WebSafe which has 216 colors, leaving plenty of room for transparency.
	p := make(color.Palette, len(palette.WebSafe)+1)
	copy(p, palette.WebSafe)
	p[len(palette.WebSafe)] = color.Transparent // Add transparent at the end

	bounds := image.Rect(0, 0, width, height)
	dst := image.NewPaletted(bounds, p)

	// Draw using FloydSteinberg or just standard Draw (NearestNeighbor)
	// Using Draw is safter for keeping sharp edges of the icon.
	draw.Draw(dst, bounds, dc.Image(), image.Point{}, draw.Src)

	return dst
}
