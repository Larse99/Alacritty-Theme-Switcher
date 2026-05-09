#!/usr/bin/env bash
set -e

REPO="https://github.com/Larse99/Alacritty-Theme-Switcher"
DIR="Alacritty-Theme-Switcher"

if ! command -v go &>/dev/null; then
    echo "Error: go is not installed" >&2
    exit 1
fi

if ! command -v git &>/dev/null; then
    echo "Error: git is not installed" >&2
    exit 1
fi

git clone "$REPO"
cd "$DIR"

echo "Building..."
go build -o ats ./cmd/

if command -v upx &>/dev/null; then
    echo "Compressing binary with UPX..."
    upx --best ats
fi

echo "Done. Binary: $(pwd)/ats"
