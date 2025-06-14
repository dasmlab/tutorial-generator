package writer

import (
	"fmt"
	"io"
	"time"

	"subtitle-exporter/parser"
)

// writeTimestamp converts a float seconds to VTT timestamp string
func writeTimestamp(seconds float64) string {
	d := time.Duration(seconds * float64(time.Second))
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	ms := int(d.Milliseconds()) % 1000
	return fmt.Sprintf("%02d:%02d:%02d.%03d", h, m, s, ms)
}

// WriteVTT emits a VTT subtitle file from parsed slide data
func WriteVTT(slides []parser.Slide, w io.Writer) error {
	_, err := fmt.Fprintln(w, "WEBVTT\n")
	if err != nil {
		return err
	}

	index := 1
	for _, slide := range slides {
		if slide.Voice == "" {
			continue
		}

		start := writeTimestamp(slide.At)
		end := writeTimestamp(slide.At + slide.Duration)

		fmt.Fprintf(w, "%d\n", index)
		fmt.Fprintf(w, "%s --> %s\n", start, end)
		fmt.Fprintf(w, "%s\n\n", slide.Voice)
		index++
	}

	return nil
}

