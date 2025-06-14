package main
// will be move to cmd/ directory eventually

import (
	"fmt"
	"os"

	"tutorial-generator/renderer"
	"tutorial-generator/player"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <manifest.yaml> <output.json>")
		os.Exit(1)
	}

	manifestFile := os.Args[1]
	outputFile := os.Args[2]

	slides, err := renderer.LoadManifest(manifestFile)
	if err != nil {
		panic(err)
	}

	timeline := player.BuildTimeline(slides)
	err = player.DumpTimelineToFile(timeline, outputFile)
	if err != nil {
		panic(err)
	}
}

