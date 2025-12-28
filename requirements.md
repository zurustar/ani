# Requirements

## Summary
A CLI tool to create an animated GIF where a specified small image moves from left to right across a transparent background.

## Functional Requirements
1.  **Input**:
    -   Path to the source image (the "small image").
    -   Duration of the animation (in seconds).
    -   Output GIF width (height will be same as input image).
    -   Delay per frame (in centiseconds, 1/100s). Default to typical value (e.g., 3-4cs for ~30fps).
2.  **Animation**:
    -   The source image starts at the left edge (or off-screen left).
    -   Moves linearly to the right edge (or off-screen right) over the specified duration.
    -   Frame rate should be sufficient for smooth movement (e.g., 30 fps).
3.  **Output**:
    -   A GIF file.
    -   **Background**: Transparent. Only the source image is visible.
4.  **CLI Interface**:
    -   Flags for input file, output file, duration, width, height.

## Non-Functional Requirements
-   Written in Go.
-   Uses `image/gif` for encoding.
-   Uses `github.com/fogleman/gg` (or standard `image/draw`) for compositing.
