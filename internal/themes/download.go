package themes

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const themesArchiveURL = "https://github.com/alacritty/alacritty-theme/archive/refs/heads/master.tar.gz"

func Download(themesDir string) (int, error) {
	if err := os.MkdirAll(themesDir, 0755); err != nil {
		return 0, fmt.Errorf("could not create themes directory: %w", err)
	}

	resp, err := http.Get(themesArchiveURL)
	if err != nil {
		return 0, fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("download failed: HTTP %d", resp.StatusCode)
	}

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("could not read archive: %w", err)
	}
	defer gz.Close()

	tr := tar.NewReader(gz)
	count := 0

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return count, fmt.Errorf("error reading archive: %w", err)
		}

		// only extract files from the themes/ subdirectory
		parts := strings.SplitN(hdr.Name, "/", 3)
		if len(parts) != 3 || parts[1] != "themes" || hdr.Typeflag != tar.TypeReg {
			continue
		}

		name := parts[2]
		if strings.ToLower(filepath.Ext(name)) != ".toml" {
			continue
		}

		dest := filepath.Join(themesDir, name)
		f, err := os.Create(dest)
		if err != nil {
			return count, fmt.Errorf("could not write %s: %w", name, err)
		}

		if _, err := io.Copy(f, tr); err != nil {
			f.Close()
			return count, fmt.Errorf("could not write %s: %w", name, err)
		}
		f.Close()
		count++
	}

	return count, nil
}
