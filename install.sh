#!/usr/bin/env bash
set -e

REPO="https://github.com/Larse99/Alacritty-Theme-Switcher"
DIR="Alacritty-Theme-Switcher"
BINARY="ats"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"

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
go build -o "$BINARY" ./cmd/

if command -v upx &>/dev/null; then
    echo "Compressing binary with UPX..."
    upx --best "$BINARY"
fi

mkdir -p "$INSTALL_DIR"
mv "$BINARY" "$INSTALL_DIR/$BINARY"
echo "✓ Installed to $INSTALL_DIR/$BINARY"

if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
    echo ""
    echo "Note: $INSTALL_DIR is not in your PATH."
    echo "Add this to your shell config (~/.bashrc, ~/.zshrc, etc.):"
    echo ""
    echo "  export PATH=\"\$HOME/.local/bin:\$PATH\""
fi
