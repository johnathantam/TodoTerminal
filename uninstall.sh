#!/usr/bin/env bash
BINARY_NAME="todo"
INSTALL_DIR="/usr/local/bin"

if [ -f "${INSTALL_DIR}/${BINARY_NAME}" ]; then
  sudo rm "${INSTALL_DIR}/${BINARY_NAME}"
  echo "${BINARY_NAME} has been uninstalled."
else
  echo "${BINARY_NAME} is not installed at ${INSTALL_DIR}/${BINARY_NAME}."
fi