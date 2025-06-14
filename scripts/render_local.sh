#!/bin/bash
set -e

MANIFEST_FILE=$1
OUTFILE="./assets/preview_timeline.json"

echo "🧾 Parsing manifest: $MANIFEST_FILE"
go run main.go "$MANIFEST_FILE" "$OUTFILE"

echo "✅ Timeline written to: $OUTFILE"

