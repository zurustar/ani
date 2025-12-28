package animator

import (
	"image"

	"image/gif"
	"math"

	"github.com/zurustar/ani/renderer"
)

// Animator handles the GIF generation logic.
type Animator struct {
	InputImage image.Image
	Duration   float64 // in seconds
	Delay      int     // in centiseconds (1/100s)
	Width      int
	Height     int
}

// NewAnimator creates a new Animator instance.
func NewAnimator(inputImage image.Image, duration float64, delay int, width int) *Animator {
	return &Animator{
		InputImage: inputImage,
		Duration:   duration,
		Delay:      delay,
		Width:      width,
		Height:     inputImage.Bounds().Dy(),
	}
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

	outGIF := &gif.GIF{
		LoopCount: 0, // Infinite loop
	}

	for i := 0; i < totalFrames; i++ {
		x := int(float64(i) * stepSize)

		// Use renderer to draw the frame
		// For now, we use Plan9 palette for simplicity as determined in design.md.
		// Transparency handling will be inside renderer.
		frame := renderer.RenderFrame(a.InputImage, a.Width, a.Height, x, 0)

		outGIF.Image = append(outGIF.Image, frame)
		outGIF.Delay = append(outGIF.Delay, a.Delay)
		// DisposalMethod 2 (Background) clears the frame, good for transparency.
		outGIF.Disposal = append(outGIF.Disposal, gif.DisposalBackground)
	}

	return outGIF, nil
}
