package fortune

import (
	"embed"
	"io/fs"
	"math/rand"
	"strings"
	
	"github.com/krry/kamadhenu/internal/config"
	"github.com/krry/kamadhenu/internal/sources"
)

var fortunes []string

// Init loads fortunes from config sources, falls back to embedded
func Init(files embed.FS) error {
	fortunes = make([]string, 0)
	
	// Load config
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	
	// If sources configured, load from them
	if len(cfg.Sources) > 0 {
		for _, source := range cfg.Sources {
			sourceFortunes, err := sources.LoadFromPath(source)
			if err != nil {
				// Log error but continue
				continue
			}
			fortunes = append(fortunes, sourceFortunes...)
		}
	}
	
	// If no fortunes loaded from sources, use embedded
	if len(fortunes) == 0 {
		err = loadEmbedded(files)
	}
	
	return err
}

func loadEmbedded(files embed.FS) error {
	// Read all fortune files (not .dat files)
	err := fs.WalkDir(files, "fortunes", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		if d.IsDir() || strings.HasSuffix(path, ".dat") {
			return nil
		}
		
		data, err := files.ReadFile(path)
		if err != nil {
			return err
		}
		
		// Split by % delimiter (fortune file format)
		entries := strings.Split(string(data), "\n%\n")
		for _, entry := range entries {
			entry = strings.TrimSpace(entry)
			if entry != "" && len(entry) <= 320 { // Match -sn 320 flag
				fortunes = append(fortunes, entry)
			}
		}
		
		return nil
	})
	
	return err
}

// Random returns a random fortune
func Random() string {
	if len(fortunes) == 0 {
		return "No fortunes available!"
	}
	return fortunes[rand.Intn(len(fortunes))]
}

// Count returns total number of loaded fortunes
func Count() int {
	return len(fortunes)
}
