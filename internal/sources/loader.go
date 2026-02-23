package sources

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// QuoteJSON is a common JSON quote format
type QuoteJSON struct {
	Quote  string `json:"quote"`
	Text   string `json:"text"`    // Alternative field
	Author string `json:"author"`
}

// LoadFromPath loads fortunes from a file or URL
func LoadFromPath(source string) ([]string, error) {
	// Check if it's a URL
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		return loadFromURL(source)
	}
	
	// Expand ~ to home directory
	if strings.HasPrefix(source, "~/") {
		home, _ := os.UserHomeDir()
		source = filepath.Join(home, source[2:])
	}
	
	return loadFromFile(source)
}

func loadFromFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	
	return parseContent(string(data), path)
}

func loadFromURL(url string) ([]string, error) {
	// Cache directory
	home, _ := os.UserHomeDir()
	cacheDir := filepath.Join(home, ".cache", "kamadhenu")
	os.MkdirAll(cacheDir, 0755)
	
	// Cache file path (hash URL for filename)
	cacheFile := filepath.Join(cacheDir, fmt.Sprintf("%x.cache", []byte(url)))
	
	// Check if cache exists and is fresh (< 24 hours)
	if info, err := os.Stat(cacheFile); err == nil {
		if time.Since(info.ModTime()) < 24*time.Hour {
			data, err := os.ReadFile(cacheFile)
			if err == nil {
				return parseContent(string(data), url)
			}
		}
	}
	
	// Fetch from URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	// Cache it
	os.WriteFile(cacheFile, data, 0644)
	
	return parseContent(string(data), url)
}

func parseContent(content string, source string) ([]string, error) {
	var fortunes []string
	
	// Try JSON array first
	var jsonQuotes []QuoteJSON
	if err := json.Unmarshal([]byte(content), &jsonQuotes); err == nil {
		for _, q := range jsonQuotes {
			text := q.Quote
			if text == "" {
				text = q.Text
			}
			if text != "" {
				if q.Author != "" {
					text = text + "\n\t\t— " + q.Author
				}
				fortunes = append(fortunes, text)
			}
		}
		return fortunes, nil
	}
	
	// Try single JSON quote
	var singleQuote QuoteJSON
	if err := json.Unmarshal([]byte(content), &singleQuote); err == nil {
		text := singleQuote.Quote
		if text == "" {
			text = singleQuote.Text
		}
		if text != "" && singleQuote.Author != "" {
			text = text + "\n\t\t— " + singleQuote.Author
		}
		if text != "" {
			return []string{text}, nil
		}
	}
	
	// Try fortune format (% delimited)
	if strings.Contains(content, "\n%") {
		parts := strings.Split(content, "\n%\n")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part != "" && len(part) <= 500 { // Max length
				fortunes = append(fortunes, part)
			}
		}
		if len(fortunes) > 0 {
			return fortunes, nil
		}
	}
	
	// Fall back to paragraph breaks (double newline)
	if strings.Contains(content, "\n\n") {
		parts := strings.Split(content, "\n\n")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part != "" && len(part) <= 500 {
				fortunes = append(fortunes, part)
			}
		}
		if len(fortunes) > 0 {
			return fortunes, nil
		}
	}
	
	// Fall back to single line per fortune
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") && len(line) <= 500 {
			fortunes = append(fortunes, line)
		}
	}
	
	if len(fortunes) == 0 {
		return nil, fmt.Errorf("no fortunes found in %s", source)
	}
	
	return fortunes, nil
}
