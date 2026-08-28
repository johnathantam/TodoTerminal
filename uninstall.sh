#!/usr/bin/env bash

# Colors (only if running in an interactive terminal)
if [ -t 1 ]; then
  GREEN='\033[0;32m'
  YELLOW='\033[0;33m'
  NC='\033[0m'
else
  GREEN='' YELLOW='' NC=''
fi

BINARY_NAME="todo"
INSTALL_DIR="/usr/local/bin"

if [ -f "${INSTALL_DIR}/${BINARY_NAME}" ]; then
  sudo rm "${INSTALL_DIR}/${BINARY_NAME}"
  echo -e "${GREEN}${BINARY_NAME} has been uninstalled.${NC}"
else
  echo -e "${YELLOW}${BINARY_NAME} is not installed at ${INSTALL_DIR}/${BINARY_NAME}.${NC}"
fi