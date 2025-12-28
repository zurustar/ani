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
-   **Palette**: 
    -   We generate a dynamic palette based on the unique colors in the input image.
    -   If the image has <= 255 unique colors, we preserve them exactly (plus transparency).
    -   If the image has > 255 unique colors, we fallback to `palette.Plan9` but ensure transparency is included.
-   **Quantization**: For V1, the simple dynamic palette approach covers most use cases (icons, simple logos). Complex images with >256 colors will use the fallback standard palette.

### Coordinate System
-   `x = 0`: Image left edge is at canvas left edge.
-   Movement: Start at `x=0` (or `x` such that image is just entering?), End at `x = width - image_width` (or `x = width`?).
-   **Decision**: Requirements say "starts at left edge... moves to right edge".
    -   Start: `x = 0` (Fully visible on left)
    -   End: `x = canvas_width - image_width` (Fully visible on right)
    -   *Clarification*: If the user wants it to fly *in* and fly *out*, we can adjust, but "move left to right" usually implies staying within bounds or traversing. Let's assume **Left-aligned to Right-aligned** traversal for V1.

### Mathematical Logic
-   **Frame Count Calculation**: `Total Frames = Floor((Duration * 100) / Delay)`
    -   *Note*: If `Duration` is not perfectly divisible by `Delay`, the actual total duration of the GIF will be slightly less than the specified `Duration`. We prioritize keeping the usage of `Delay` consistent for generic GIF players.
-   `Step Size (pixels/frame) = (CanvasWidth - ImageWidth) / (TotalFrames - 1)`
-   Frame `i`: `x = i * StepSize`

## Libraries
-   Standard `flag` for CLI.
-   `image/gif`, `image/png` for I/O.
-   `github.com/fogleman/gg` for easy drawing.
