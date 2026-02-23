#!/usr/bin/env bash
# Quick install script for Kamadhenu

set -e

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

case "$OS" in
    darwin) OS="darwin" ;;
    linux) OS="linux" ;;
    *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

BINARY="kamadhenu-${OS}-${ARCH}"
BASE_URL="https://github.com/krry/Kamadhenu/releases/latest/download"

echo "Installing Kamadhenu for ${OS}/${ARCH}..."

# Download binary
echo "Downloading binary..."
curl -fsSL "${BASE_URL}/${BINARY}" -o kamadhenu
chmod +x kamadhenu

# Download man page
echo "Downloading man page..."
curl -fsSL "${BASE_URL}/kamadhenu.1" -o kamadhenu.1

# Install
echo "Installing to /usr/local/bin (requires sudo)..."
sudo install -m 755 kamadhenu /usr/local/bin/
sudo mkdir -p /usr/local/share/man/man1
sudo install -m 644 kamadhenu.1 /usr/local/share/man/man1/

# Cleanup
rm -f kamadhenu kamadhenu.1

echo "✅ Kamadhenu installed successfully!"
echo ""
echo "Try it:"
echo "  kamadhenu"
echo "  kamadhenu 'Hello, world'"
echo "  man kamadhenu"
echo ""
echo "Tip: Add 'alias kama=kamadhenu' to your shell config for quick access."
