package main

import (
	"flag"
	"fmt"
	"image/gif"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"image/color"

	"github.com/zurustar/ani/animator"
	"github.com/zurustar/ani/renderer"
)

func main() {
	inputPath := flag.String("i", "", "Path to the source image (PNG only)")
	outputPath := flag.String("o", "output.gif", "Path to the output GIF file")
	duration := flag.Float64("duration", 60.0, "Duration of the animation in seconds")
	width := flag.Int("width", 0, "Width of the output GIF")

	delay := flag.Int("delay", 4, "Delay per frame in centiseconds (1/100s)")
	bgHex := flag.String("bg", "", "Background color in #RRGGBB format (default transparent)")

	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Printf("Error: Unexpected arguments found: %v\n", flag.Args())
		fmt.Println("Did you provide a value without a flag or forget a '-' prefix?")
		flag.Usage()
		os.Exit(1)
	}

	if *inputPath == "" {
		fmt.Println("Error: Input file is required (-i)")
		flag.Usage()
		os.Exit(1)
	}

	if *width <= 0 {
		fmt.Println("Error: Width must be positive (-width)")
		flag.Usage()
		os.Exit(1)
	}

	// 1. Load Image
	file, err := os.Open(*inputPath)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(*inputPath))
	if ext != ".png" {
		fmt.Println("Error: Only PNG format is supported")
		os.Exit(1)
	}

	img, err := png.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding PNG: %v\n", err)
		os.Exit(1)
	}

	// 1.5 Parse Background Color
	var bgColor color.Color = color.Transparent
	if *bgHex != "" {
		c, err := renderer.ParseHexColor(*bgHex)
		if err != nil {
			fmt.Printf("Error parsing background color: %v\n", err)
			flag.Usage()
			os.Exit(1)
		}
		bgColor = c
	}

	// 2. Initialize Animator
	a := animator.NewAnimator(img, *duration, *delay, *width, bgColor)

	// 3. Generate GIF
	g, err := a.GenerateGIF()
	if err != nil {
		fmt.Printf("Error generating GIF: %v\n", err)
		os.Exit(1)
	}

	// 4. Save GIF
	outFile, err := os.Create(*outputPath)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	err = gif.EncodeAll(outFile, g)
	if err != nil {
		fmt.Printf("Error encoding GIF: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated text animation GIF: %s\n", *outputPath)
}
