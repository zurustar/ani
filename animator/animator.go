package animator

import (
	"image"

	"image/color"
	"image/color/palette"
	"image/gif"
	"math"

	"github.com/zurustar/ani/renderer"
)

// CalculateTotalFrames calculates the number of frames based on duration and delay.
func CalculateTotalFrames(duration float64, delay int) int {
	totalFrames := int(math.Floor((duration * 100) / float64(delay)))
	if totalFrames <= 0 {
		return 1
	}
	return totalFrames
}

// CalculateStepSize calculates the pixels to move per frame.
func CalculateStepSize(canvasWidth, imgWidth, totalFrames int) float64 {
	if totalFrames <= 1 {
		return 0
	}
	return float64(canvasWidth-imgWidth) / float64(totalFrames-1)
}

// Animator handles the GIF generation logic.
type Animator struct {
	InputImage      image.Image
	Duration        float64 // in seconds
	Delay           int     // in centiseconds (1/100s)
	Width           int
	Height          int
	BackgroundColor color.Color
}

// NewAnimator creates a new Animator instance.
func NewAnimator(inputImage image.Image, duration float64, delay int, width int, bgColor color.Color) *Animator {
	return &Animator{
		InputImage:      inputImage,
		Duration:        duration,
		Delay:           delay,
		Width:           width,
		Height:          inputImage.Bounds().Dy(),
		BackgroundColor: bgColor,
	}
}

// generatePalette inspects the image and creates an optimal palette.
func generatePalette(img image.Image, bgColor color.Color) color.Palette {
	// Simple color frequency map
	colors := make(map[color.Color]bool)
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			colors[img.At(x, y)] = true
		}
	}

	// Add background color to palette if provided
	if bgColor != nil && bgColor != color.Transparent {
		colors[bgColor] = true
	}

	// If total colors <= 255, we can preserve them exactly (leaving room for transparency)
	if len(colors) <= 255 {
		p := make(color.Palette, 0, len(colors)+1)
		for c := range colors {
			p = append(p, c)
		}
		p = append(p, color.Transparent)
		return p
	}

	// Fallback to Plan9 + transparent if too many colors
	p := make(color.Palette, len(palette.Plan9)+1)
	copy(p, palette.Plan9)
	p[len(palette.Plan9)] = color.Transparent
	return p
}

// GenerateGIF creates the animated GIF.
func (a *Animator) GenerateGIF() (*gif.GIF, error) {
	// Calculate total frames: Floor((Duration * 100) / Delay)
	totalFrames := int(math.Floor((a.Duration * 100) / float64(a.Delay)))
	if totalFrames <= 0 {
		totalFrames = 1
	}

	imgWidth := a.InputImage.Bounds().Dx()
	// Step Size = (CanvasWidth - ImageWidth) / (TotalFrames - 1)
	// If TotalFrames is 1, stepSize is 0.
	var stepSize float64
	if totalFrames > 1 {
		stepSize = float64(a.Width-imgWidth) / float64(totalFrames-1)
	}

	// Generate palette once
	p := generatePalette(a.InputImage, a.BackgroundColor)

	outGIF := &gif.GIF{
		LoopCount: 0, // Infinite loop
	}

	for i := 0; i < totalFrames; i++ {
		x := int(float64(i) * stepSize)

		// Use renderer to draw the frame
		// Pass the custom palette
		frame := renderer.RenderFrame(a.InputImage, a.Width, a.Height, x, 0, p, a.BackgroundColor)

		outGIF.Image = append(outGIF.Image, frame)
		outGIF.Delay = append(outGIF.Delay, a.Delay)
		// DisposalMethod 2 (Background) clears the frame, good for transparency.
		outGIF.Disposal = append(outGIF.Disposal, gif.DisposalBackground)
	}

	return outGIF, nil
}
