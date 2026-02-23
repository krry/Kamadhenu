package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Sources []string `yaml:"sources"` // File paths or URLs
}

var configPath string

func init() {
	home, _ := os.UserHomeDir()
	configDir := filepath.Join(home, ".config", "kamadhenu")
	os.MkdirAll(configDir, 0755)
	configPath = filepath.Join(configDir, "config.yaml")
}

// Load reads the config file, returns empty config if file doesn't exist
func Load() (*Config, error) {
	cfg := &Config{
		Sources: []string{},
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil // No config yet, return empty
		}
		return nil, err
	}

	err = yaml.Unmarshal(data, cfg)
	return cfg, err
}

// Save writes the config to disk
func (c *Config) Save() error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}

// AddSource adds a source if it doesn't already exist
func (c *Config) AddSource(source string) bool {
	for _, s := range c.Sources {
		if s == source {
			return false // Already exists
		}
	}
	c.Sources = append(c.Sources, source)
	return true
}

// RemoveSource removes a source, returns true if found
func (c *Config) RemoveSource(source string) bool {
	for i, s := range c.Sources {
		if s == source {
			c.Sources = append(c.Sources[:i], c.Sources[i+1:]...)
			return true
		}
	}
	return false
}

// Reset clears all sources
func (c *Config) Reset() {
	c.Sources = []string{}
}

// ConfigPath returns the path to the config file
func ConfigPath() string {
	return configPath
}
