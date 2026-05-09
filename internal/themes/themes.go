package themes

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func List(themesDir string) ([]string, error) {
	entries, err := os.ReadDir(themesDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("themes directory not found: %s", themesDir)
		}
		return nil, fmt.Errorf("could not read themes directory: %w", err)
	}

	var names []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if strings.ToLower(filepath.Ext(e.Name())) == ".toml" {
			names = append(names, strings.TrimSuffix(e.Name(), filepath.Ext(e.Name())))
		}
	}

	sort.Strings(names)
	return names, nil
}

func Switch(themesDir, currentThemePath, alacrittyConfigPath, name string) error {
	themePath := filepath.Join(themesDir, name+".toml")

	data, err := os.ReadFile(themePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("theme '%s' not found", name)
	}
	if err != nil {
		return fmt.Errorf("permission denied reading theme '%s'", name)
	}

	content := append([]byte("# ats:theme="+name+"\n"), data...)
	if err := os.WriteFile(currentThemePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write theme: %w", err)
	}

	now := time.Now()
	if err := os.Chtimes(alacrittyConfigPath, now, now); err != nil {
		return fmt.Errorf("theme written but could not trigger reload (check alacritty_config path): %w", err)
	}

	return nil
}

func Current(currentThemePath string) (string, error) {
	f, err := os.Open(currentThemePath)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("no active theme set")
	}
	if err != nil {
		return "", fmt.Errorf("could not determine current theme")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		line := scanner.Text()
		name := strings.TrimPrefix(line, "# ats:theme=")
		if name != line {
			return name, nil
		}
	}
	return "", fmt.Errorf("no active theme set")
}
