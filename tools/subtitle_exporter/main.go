package main

import (
	"fmt"
	"os"

	"subtitle-exporter/parser"
	"subtitle-exporter/writer"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <tutorial.yaml>")
		os.Exit(1)
	}

	manifestPath := os.Args[1]

	slides, err := parser.LoadSlides(manifestPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing manifest: %v\n", err)
		os.Exit(1)
	}

	err = writer.WriteVTT(slides, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing VTT: %v\n", err)
		os.Exit(1)
	}
}

