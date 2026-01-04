package renderer

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/fogleman/gg"
)

// ParseHexColor parses a hex string (e.g., "#RRGGBB") into color.Color.
// It enforces strict 6-digit hex with a leading '#'.
func ParseHexColor(s string) (color.Color, error) {
	if len(s) != 7 || s[0] != '#' {
		return nil, fmt.Errorf("invalid color format: must be #RRGGBB")
	}

	hexToByte := func(b byte) (byte, error) {
		switch {
		case b >= '0' && b <= '9':
			return b - '0', nil
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10, nil
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10, nil
		}
		return 0, fmt.Errorf("invalid hex character: %c", b)
	}

	var rgb [3]byte
	for i := 0; i < 3; i++ {
		high, err := hexToByte(s[1+i*2])
		if err != nil {
			return nil, err
		}
		low, err := hexToByte(s[1+i*2+1])
		if err != nil {
			return nil, err
		}
		rgb[i] = (high << 4) | low
	}

	return color.RGBA{R: rgb[0], G: rgb[1], B: rgb[2], A: 255}, nil
}

// RenderFrame draws a single frame with the image at (x, y).
// It returns a paletted image with transparency, using the provided palette.
func RenderFrame(img image.Image, width, height, x, y int, p color.Palette, bgColor color.Color) *image.Paletted {
	dc := gg.NewContext(width, height)

	// Create a background
	if bgColor == nil {
		dc.SetColor(color.Transparent)
	} else {
		dc.SetColor(bgColor)
	}
	dc.Clear()

	// Draw the image
	dc.DrawImage(img, x, y)

	bounds := image.Rect(0, 0, width, height)
	dst := image.NewPaletted(bounds, p)

	// Draw using draw.Src to preserve exact colors where possible
	draw.Draw(dst, bounds, dc.Image(), image.Point{}, draw.Src)

	return dst
}
