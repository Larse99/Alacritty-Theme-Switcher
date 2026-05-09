package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ats/internal/config"
	"ats/internal/themes"
)

func main() {
	selectTheme := flag.String("s", "", "Switch to a theme by name")
	listThemes := flag.Bool("l", false, "List all available themes")
	currentTheme := flag.Bool("c", false, "Show the currently active theme")
	downloadThemes := flag.Bool("d", false, "Download themes from alacritty/alacritty-theme")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: ats [option]")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Options:")
		fmt.Fprintln(os.Stderr, "  -s <name>   Switch to a theme")
		fmt.Fprintln(os.Stderr, "  -l          List all available themes")
		fmt.Fprintln(os.Stderr, "  -c          Show the currently active theme")
		fmt.Fprintln(os.Stderr, "  -d          Download themes from alacritty/alacritty-theme")
	}

	flag.Parse()

	if !*listThemes && !*currentTheme && !*downloadThemes && *selectTheme == "" {
		flag.Usage()
		os.Exit(1)
	}

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	switch {
	case *downloadThemes:
		fmt.Println("Downloading themes...")
		count, err := themes.Download(cfg.Paths.ThemesDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("✓ Downloaded %d themes to %s\n", count, cfg.Paths.ThemesDir)

	case *listThemes:
		names, err := themes.List(cfg.Paths.ThemesDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(strings.Join(names, "\n"))

	case *selectTheme != "":
		if err := themes.Switch(cfg.Paths.ThemesDir, cfg.Paths.CurrentTheme, cfg.Paths.AlacrittyConfig, *selectTheme); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("✓ Switched theme to %s\n", *selectTheme)

	case *currentTheme:
		name, err := themes.Current(cfg.Paths.CurrentTheme)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Current theme: %s\n", name)
	}
}
