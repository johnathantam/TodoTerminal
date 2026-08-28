#!/usr/bin/env bash
set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
BOLD='\033[1m'
NC='\033[0m' # No Color, resets formatting

REPO="johnathantam/TodoTerminal"
BINARY_NAME="todo"
INSTALL_DIR="/usr/local/bin"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
  x86_64) ARCH="x86_64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo -e "${RED}Unsupported architecture: $ARCH${NC}"; exit 1 ;;
esac

case "$OS" in
  darwin) OS_NAME="Darwin" ;;
  linux) OS_NAME="Linux" ;;
  *) echo -e "${RED}Unsupported OS: $OS${NC}"; exit 1 ;;
esac

LATEST_TAG=$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')
ARCHIVE="TodoTerminal_${OS_NAME}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/${ARCHIVE}"

echo -e "${BLUE}Downloading ${BINARY_NAME} ${LATEST_TAG} for ${OS_NAME}/${ARCH}...${NC}"
TMP_DIR=$(mktemp -d)
curl -fsSL "$URL" -o "${TMP_DIR}/${ARCHIVE}"
tar -xzf "${TMP_DIR}/${ARCHIVE}" -C "$TMP_DIR"

echo -e "${BLUE}Installing to ${INSTALL_DIR}/${BINARY_NAME} (may prompt for sudo)...${NC}"
sudo mkdir -p "$INSTALL_DIR"
sudo mv "${TMP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"
sudo chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
rm -rf "$TMP_DIR"

echo -e "${GREEN}${BOLD}Installed!${NC} Run '${BINARY_NAME} help' to get started."