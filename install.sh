#!/usr/bin/env bash
set -e

REPO="johnathantam/TodoTerminal"
BINARY_NAME="todo"
INSTALL_DIR="/usr/local/bin"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
  x86_64) ARCH="x86_64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

case "$OS" in
  darwin) OS_NAME="Darwin" ;;
  linux) OS_NAME="Linux" ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

LATEST_TAG=$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')
ARCHIVE="TodoTerminal_${OS_NAME}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/${ARCHIVE}"

echo "Downloading ${BINARY_NAME} ${LATEST_TAG} for ${OS_NAME}/${ARCH}..."
TMP_DIR=$(mktemp -d)
curl -fsSL "$URL" -o "${TMP_DIR}/${ARCHIVE}"
tar -xzf "${TMP_DIR}/${ARCHIVE}" -C "$TMP_DIR"

echo "Installing to ${INSTALL_DIR}/${BINARY_NAME} (may prompt for sudo)..."
sudo mv "${TMP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"
sudo chmod +x "${INSTALL_DIR}/${BINARY_NAME}"

rm -rf "$TMP_DIR"
echo "Installed! Run '${BINARY_NAME} help' to get started."