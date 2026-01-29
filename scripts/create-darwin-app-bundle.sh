#!/usr/bin/env bash

set -euo pipefail

OS="$1"
ARCH="$2"
BINARY_PATH="$3"
APP_NAME="psx-memcard"
APP_BUNDLE_DIR="dist/${APP_NAME}.app"

if [ "$OS" != "darwin" ]; then
  echo "Skip creating app bundle for non darwin $OS."
  exit 0
fi

# Create the app bundle directory structure
mkdir -p "${APP_BUNDLE_DIR}"
cp -r build/darwin/. ${APP_BUNDLE_DIR}/

# Copy the binary into the app bundle
cp "$BINARY_PATH" "${APP_BUNDLE_DIR}/Contents/MacOS/${APP_NAME}"
rm "${APP_BUNDLE_DIR}/Contents/MacOS/.gitkeep"

(cd dist && zip -r "${APP_NAME}-${OS}-${ARCH}.app.zip" "${APP_NAME}.app")
rm -rf "${APP_BUNDLE_DIR}"
