package animator

import (
	"math"
	"testing"
)

func TestCalculateTotalFrames(t *testing.T) {
	tests := []struct {
		name     string
		duration float64
		delay    int
		want     int
	}{
		{"1 sec, 10cs delay", 1.0, 10, 10},
		{"2 sec, 4cs delay", 2.0, 4, 50},
		{"0.5 sec, 20cs delay", 0.5, 20, 2},
		{"Zero Duration", 0.0, 10, 1}, // Minimum 1 frame
		{"Small Duration", 0.01, 10, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateTotalFrames(tt.duration, tt.delay)
			if got != tt.want {
				t.Errorf("CalculateTotalFrames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateStepSize(t *testing.T) {
	tests := []struct {
		name        string
		canvasWidth int
		imgWidth    int
		totalFrames int
		want        float64
	}{
		{"Simple Move", 100, 10, 10, 10.0},       // (100-10) / 9 = 10
		{"Static Image", 100, 10, 1, 0.0},        // 1 frame -> no step
		{"No Movement Space", 100, 100, 10, 0.0}, // (100-100) / 9 = 0
		{"Small Step", 20, 10, 5, 2.5},           // (20-10) / 4 = 2.5
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateStepSize(tt.canvasWidth, tt.imgWidth, tt.totalFrames)
			if math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("CalculateStepSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
