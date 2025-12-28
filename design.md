# Design Document

## Architecture

The application will be a single Command Line Interface (CLI) tool.

### Components
1.  **CLI Parser (`main` pkg)**:
    -   Parses command line flags (`-i`, `-o`, `-duration`, `-width`, `-delay`).
    -   Validates inputs (file existence, positive dimensions/duration).
2.  **Animator (`animator` pkg)**:
    -   Responsible for generating the `gif.GIF` struct.
    -   Calculates the number of frames based on `duration` and `delay`.
    -   Calculates the X-coordinate of the image for each frame (linear interpolation).
3.  **Frame Renderer (`renderer` pkg)**:
    -   Uses `github.com/fogleman/gg` to draw each frame.
    -   Creates a transparent context of size `width x input_height`.
    -   Draws the input image at the calculated (x, 0) position.
    -   Converts the drawn `image.Image` to `*image.Paletted` for GIF encoding.

## Design Decisions

### GIF Optimization & Palette
-   **Transparency**: The background will be transparent. We need to ensure the `image/gif` encoder handles the transparent index correctly.
-   **Palette**: We will use a standard palette (e.g., `palette.Plan9` or generated from the image) but ensure index 0 (or specific index) is reserved for transparency.
-   **Quantization**: For simplicity in V1, we will use a standard palette. If the input image has many colors, we might need a quantizer, but starting with a standard palette is safer for "simple/small images".

### Coordinate System
-   `x = 0`: Image left edge is at canvas left edge.
-   Movement: Start at `x=0` (or `x` such that image is just entering?), End at `x = width - image_width` (or `x = width`?).
-   **Decision**: Requirements say "starts at left edge... moves to right edge".
    -   Start: `x = 0` (Fully visible on left)
    -   End: `x = canvas_width - image_width` (Fully visible on right)
    -   *Clarification*: If the user wants it to fly *in* and fly *out*, we can adjust, but "move left to right" usually implies staying within bounds or traversing. Let's assume **Left-aligned to Right-aligned** traversal for V1.

### Mathematical Logic
-   `Total Frames = (Duration * 100) / Delay`
-   `Step Size (pixels/frame) = (CanvasWidth - ImageWidth) / (TotalFrames - 1)`
-   Frame `i`: `x = i * StepSize`

## Libraries
-   Standard `flag` for CLI.
-   `image/gif`, `image/png`, `image/jpeg` for I/O.
-   `github.com/fogleman/gg` for easy drawing.
