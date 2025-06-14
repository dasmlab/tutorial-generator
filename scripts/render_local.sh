#!/bin/bash

set -e

TUTORIAL_PATH="$1"

if [ -z "$TUTORIAL_PATH" ]; then
  echo "Usage: ./render_local.sh <path/to/tutorial.yaml>"
  exit 1
fi

echo "🔍 Parsing manifest..."
go run renderer/manifest_loader.go "$TUTORIAL_PATH"

echo "🎬 Simulating playback..."
go run player/engine.go

