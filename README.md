# ats — Alacritty Theme Switcher

A fast, minimal CLI tool to switch Alacritty themes from the terminal.

## Usage

```bash
ats -l              # list all available themes
ats -s tokyo_night  # switch to a theme
ats -c              # show the currently active theme
```

## Installation

**Build from source:**

```bash
git clone https://github.com/leissink/Alacritty-Theme-Switcher
cd Alacritty-Theme-Switcher
go build -o ats ./cmd/
```

Then move the binary somewhere on your `$PATH`:

```bash
mv ats ~/.local/bin/
```

**Or install directly with go install:**

```bash
go install ats@latest
```

## Setup

### 1. Alacritty config

Add the following import to your Alacritty config (`~/.config/alacritty/.alacritty.toml`):

```toml
[general]
import = [
  "~/.config/alacritty/current-theme.toml"
]
```

`ats` will create and update `current-theme.toml` as a symlink pointing to the active theme file.

### 2. Themes directory

Place your `.toml` theme files in:

```
~/.config/alacritty/themes/
```

### 3. Config file

On first run, `ats` creates a config file at `~/.config/ats/config.toml` with these defaults:

```toml
[paths]
alacritty_config = "~/.config/alacritty/.alacritty.toml"
current_theme = "~/.config/alacritty/current-theme.toml"
themes_dir = "~/.config/alacritty/themes"
```

Adjust paths as needed. Both `~` and environment variables (e.g. `$HOME`) are supported.

## Requirements

- Go 1.22+
- macOS, Linux (Fedora, Ubuntu)
