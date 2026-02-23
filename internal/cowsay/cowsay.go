package cowsay

import (
	"embed"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"golang.org/x/term"
)

var cows map[string]string

// Init loads all cow files from embedded FS
func Init(files embed.FS) error {
	cows = make(map[string]string)
	
	err := fs.WalkDir(files, "cows", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		if d.IsDir() || !strings.HasSuffix(path, ".cow") {
			return nil
		}
		
		data, err := files.ReadFile(path)
		if err != nil {
			return err
		}
		
		name := filepath.Base(path)
		name = strings.TrimSuffix(name, ".cow")
		cows[name] = string(data)
		
		return nil
	})
	
	return err
}

// Random returns a random cow name
func Random() string {
	keys := make([]string, 0, len(cows))
	for k := range cows {
		keys = append(keys, k)
	}
	if len(keys) == 0 {
		return "default"
	}
	return keys[rand.Intn(len(keys))]
}

// Say generates cowsay output with word wrapping
func Say(text, cowName string) string {
	cow, exists := cows[cowName]
	if !exists {
		cow = getDefaultCow()
	}
	
	// Get terminal width
	width := 80
	if w, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil && w > 0 {
		width = w - 3 // Leave margin
	}
	if width > 128 {
		width = 128
	}
	
	// Wrap text
	lines := wordWrap(text, width-4) // Leave room for bubble borders
	
	// Build bubble
	bubble := makeBubble(lines)
	
	// Parse and render cow
	cowArt := renderCow(cow)
	
	return bubble + "\n" + cowArt
}

func wordWrap(text string, width int) []string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{""}
	}
	
	lines := []string{}
	currentLine := ""
	
	for _, word := range words {
		if len(currentLine)+len(word)+1 <= width {
			if currentLine == "" {
				currentLine = word
			} else {
				currentLine += " " + word
			}
		} else {
			if currentLine != "" {
				lines = append(lines, currentLine)
			}
			currentLine = word
		}
	}
	
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	
	return lines
}

func makeBubble(lines []string) string {
	if len(lines) == 0 {
		return ""
	}
	
	// Find max line length
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	
	// Build bubble
	result := " " + strings.Repeat("_", maxLen+2) + "\n"
	
	if len(lines) == 1 {
		result += fmt.Sprintf("< %s >\n", pad(lines[0], maxLen))
	} else {
		for i, line := range lines {
			if i == 0 {
				result += fmt.Sprintf("/ %s \\\n", pad(line, maxLen))
			} else if i == len(lines)-1 {
				result += fmt.Sprintf("\\ %s /\n", pad(line, maxLen))
			} else {
				result += fmt.Sprintf("| %s |\n", pad(line, maxLen))
			}
		}
	}
	
	result += " " + strings.Repeat("-", maxLen+2)
	
	return result
}

func pad(s string, length int) string {
	if len(s) >= length {
		return s
	}
	return s + strings.Repeat(" ", length-len(s))
}

func renderCow(cowTemplate string) string {
	// Simple cow rendering - replace $thoughts and $eyes
	// Default eyes and thoughts
	eyes := "oo"
	thoughts := "\\"
	
	cowTemplate = strings.ReplaceAll(cowTemplate, "$thoughts", thoughts)
	cowTemplate = strings.ReplaceAll(cowTemplate, "$eyes", eyes)
	cowTemplate = strings.ReplaceAll(cowTemplate, "${eyes}", eyes)
	cowTemplate = strings.ReplaceAll(cowTemplate, "${thoughts}", thoughts)
	
	// Remove perl code blocks (between $EOC markers or ##)
	cowTemplate = regexp.MustCompile(`(?s)\$the_cow = <<.*?EOC;`).ReplaceAllString(cowTemplate, "")
	cowTemplate = regexp.MustCompile(`(?m)^##.*$`).ReplaceAllString(cowTemplate, "")
	
	// Extract just the ASCII art (lines that start with whitespace or special chars)
	lines := strings.Split(cowTemplate, "\n")
	art := []string{}
	inArt := false
	
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		
		// Detect art lines (start with whitespace or common art chars)
		if len(line) > 0 && (line[0] == ' ' || line[0] == '\t' || 
			strings.ContainsAny(string(line[0]), `\/|_-^().,'";:<>[]{}@#$%&*+=~` + "`")) {
			inArt = true
			art = append(art, line)
		} else if inArt {
			break
		}
	}
	
	return strings.Join(art, "\n")
}

func getDefaultCow() string {
	return `$the_cow = <<"EOC";
        $thoughts   ^__^
         $thoughts  ($eyes)\\_______
            (__)\\       )\\/\\
                ||----w |
                ||     ||
EOC
`
}

// Count returns total number of loaded cows
func Count() int {
	return len(cows)
}
