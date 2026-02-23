package colors

import (
	"fmt"
	"math"
	"math/rand"
)

type RainbowMode int

const (
	ModeNormal RainbowMode = iota
	ModeUnicorn
)

var (
	currentMode     = ModeNormal
	startingHue     float64
	phaseR, phaseG, phaseB float64
)

func init() {
	// Initialize random starting point on first use
	startingHue = rand.Float64() * 2 * math.Pi
	phaseR = 0
	phaseG = 2 * math.Pi / 3
	phaseB = 4 * math.Pi / 3
}

// SetMode configures the rainbow color mode
func SetMode(mode RainbowMode) {
	currentMode = mode
	startingHue = rand.Float64() * 2 * math.Pi
	
	if mode == ModeUnicorn {
		// Random phase shifts for chaotic color combinations
		phaseR = rand.Float64() * 2 * math.Pi
		phaseG = rand.Float64() * 2 * math.Pi
		phaseB = rand.Float64() * 2 * math.Pi
	} else {
		// Standard 120° phase shifts
		phaseR = 0
		phaseG = 2 * math.Pi / 3
		phaseB = 4 * math.Pi / 3
	}
}

// Rainbow prints text with rainbow gradient colors
func Rainbow(text string) {
	freq := 0.125
	offset := startingHue
	
	for _, char := range text {
		if char == '\n' {
			fmt.Print("\n")
			continue
		}
		
		// Calculate RGB with mode-specific phases
		r := int(math.Sin(freq*offset+phaseR)*127 + 128)
		g := int(math.Sin(freq*offset+phaseG)*127 + 128)
		b := int(math.Sin(freq*offset+phaseB)*127 + 128)
		
		// In unicorn mode, boost dark colors for readability
		if currentMode == ModeUnicorn {
			// Calculate luminance (perceived brightness)
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			
			// If too dark (< 80), boost to minimum brightness
			if lum < 80 {
				boost := 80 / lum
				r = clamp(int(float64(r)*boost), 80, 255)
				g = clamp(int(float64(g)*boost), 80, 255)
				b = clamp(int(float64(b)*boost), 80, 255)
			}
		}
		
		// Clamp values
		r = clamp(r, 0, 255)
		g = clamp(g, 0, 255)
		b = clamp(b, 0, 255)
		
		// Print with ANSI RGB color
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, char)
		
		offset += 1
	}
}

// RainbowString returns rainbow-colored version of string
func RainbowString(text string) string {
	freq := 0.125
	offset := startingHue
	result := ""
	
	for _, char := range text {
		if char == '\n' {
			result += "\n"
			continue
		}
		
		r := int(math.Sin(freq*offset+phaseR)*127 + 128)
		g := int(math.Sin(freq*offset+phaseG)*127 + 128)
		b := int(math.Sin(freq*offset+phaseB)*127 + 128)
		
		// In unicorn mode, boost dark colors for readability
		if currentMode == ModeUnicorn {
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			if lum < 80 {
				boost := 80 / lum
				r = clamp(int(float64(r)*boost), 80, 255)
				g = clamp(int(float64(g)*boost), 80, 255)
				b = clamp(int(float64(b)*boost), 80, 255)
			}
		}
		
		r = clamp(r, 0, 255)
		g = clamp(g, 0, 255)
		b = clamp(b, 0, 255)
		
		result += fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, char)
		offset += 1
	}
	
	return result
}

func clamp(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
