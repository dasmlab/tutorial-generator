#!/bin/bash
set -e
echo "🧠 Generating timeline from $1"
go run renderer/manifest_loader.go "$1" > assets/preview_timeline.json

echo "🗣️  Writing subtitle file..."
go run tts_engine/subtitle_writer.go "$1"

echo "🎬 Launching browser"
xdg-open player/index.html

