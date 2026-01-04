package renderer

import (
	"image/color"
	"testing"
)

func TestParseHexColor(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    color.RGBA
		wantErr bool
	}{
		{
			name:  "Valid Red",
			input: "#FF0000",
			want:  color.RGBA{R: 255, G: 0, B: 0, A: 255},
		},
		{
			name:  "Valid Green",
			input: "#00FF00",
			want:  color.RGBA{R: 0, G: 255, B: 0, A: 255},
		},
		{
			name:  "Valid Blue",
			input: "#0000FF",
			want:  color.RGBA{R: 0, G: 0, B: 255, A: 255},
		},
		{
			name:  "Valid Black",
			input: "#000000",
			want:  color.RGBA{R: 0, G: 0, B: 0, A: 255},
		},
		{
			name:  "Valid White",
			input: "#FFFFFF",
			want:  color.RGBA{R: 255, G: 255, B: 255, A: 255},
		},
		{
			name:    "Missing Hash",
			input:   "FF0000",
			wantErr: true,
		},
		{
			name:    "Invalid Length Short",
			input:   "#F00",
			wantErr: true,
		},
		{
			name:    "Invalid Length Long",
			input:   "#FF00000",
			wantErr: true,
		},
		{
			name:    "Invalid Characters",
			input:   "#GG0000",
			wantErr: true,
		},
		{
			name:    "Empty",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseHexColor(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseHexColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Convert to RGBA to compare
				r, g, b, a := got.RGBA()
				wr, wg, wb, wa := tt.want.RGBA()
				if r != wr || g != wg || b != wb || a != wa {
					t.Errorf("ParseHexColor() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
