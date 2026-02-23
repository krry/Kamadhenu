package typewriter

import (
	"fmt"
	"strings"
	"time"
	
	"github.com/krry/kamadhenu/internal/colors"
)

// Print outputs text line-by-line with a delay (typewriter effect)
func Print(text string, noColor bool) {
	lines := strings.Split(text, "\n")
	
	for _, line := range lines {
		if noColor {
			fmt.Println(line)
		} else {
			colors.Rainbow(line + "\n")
		}
		time.Sleep(30 * time.Millisecond) // 30ms delay between lines
	}
}
