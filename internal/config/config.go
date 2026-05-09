package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Paths struct {
		AlacrittyConfig string `toml:"alacritty_config"`
		CurrentTheme    string `toml:"current_theme"`
		ThemesDir       string `toml:"themes_dir"`
	} `toml:"paths"`
}

func Load() (*Config, error) {
	configPath, err := configFilePath()
	if err != nil {
		return nil, fmt.Errorf("could not resolve config path: %w", err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		cfg := defaults()
		if err := writeDefaults(configPath, cfg); err != nil {
			return nil, fmt.Errorf("could not create default config: %w", err)
		}
		cfg.Paths.AlacrittyConfig = expandPath(cfg.Paths.AlacrittyConfig)
		cfg.Paths.CurrentTheme = expandPath(cfg.Paths.CurrentTheme)
		cfg.Paths.ThemesDir = expandPath(cfg.Paths.ThemesDir)
		return &cfg, nil
	}

	var cfg Config
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("could not parse config %s: %w", configPath, err)
	}

	cfg.Paths.AlacrittyConfig = expandPath(cfg.Paths.AlacrittyConfig)
	cfg.Paths.CurrentTheme = expandPath(cfg.Paths.CurrentTheme)
	cfg.Paths.ThemesDir = expandPath(cfg.Paths.ThemesDir)

	return &cfg, nil
}

func configFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "ats", "config.toml"), nil
}

func defaults() Config {
	var cfg Config
	cfg.Paths.AlacrittyConfig = "~/.config/alacritty/alacritty.toml"
	cfg.Paths.CurrentTheme = "~/.config/alacritty/current-theme.toml"
	cfg.Paths.ThemesDir = "~/.config/alacritty/themes"
	return cfg
}

func writeDefaults(path string, cfg Config) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	content := fmt.Sprintf("[paths]\nalacritty_config = %q\ncurrent_theme = %q\nthemes_dir = %q\n",
		cfg.Paths.AlacrittyConfig,
		cfg.Paths.CurrentTheme,
		cfg.Paths.ThemesDir,
	)
	return os.WriteFile(path, []byte(content), 0644)
}

func expandPath(path string) string {
	path = os.ExpandEnv(path)
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err == nil {
			path = filepath.Join(home, path[2:])
		}
	}
	return path
}
