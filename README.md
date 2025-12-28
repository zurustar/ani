# ani

`ani` is a CLI tool written in Go to create simple animated GIFs.
It generates an animation where a specified source image moves horizontally from left to right across a transparent background.

## Features
-   **Horizontal Movement**: Moves a small image from left to right.
-   **Transparency**: Output GIF has a transparent background.
-   **Customizable**: Control duration, width, and frame delay.

## Installation

```bash
go install github.com/zurustar/ani@latest
```

## Usage

```bash
ani -i <input_image> [flags]
```

### Flags
-   `-i string`: Path to the source image (PNG only) [Required]
-   `-o string`: Path to the output GIF file (default "output.gif")
-   `-duration float`: Duration of the animation in seconds (default 1.0)
-   `-width int`: Width of the output GIF [Required]
-   `-delay int`: Delay per frame in centiseconds (default 4, approx 25fps)

### Example

```bash
# Create a 300px wide GIF where 'icon.png' moves across in 2 seconds
ani -i icon.png -o animation.gif -width 300 -duration 2.0
```