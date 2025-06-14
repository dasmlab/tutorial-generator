package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "time"

    "gopkg.in/yaml.v3"
)

type Slide struct {
    Title       string `yaml:"title"`
    Description string `yaml:"description"`
    Voice       string `yaml:"voice"`
    At          int    `yaml:"at"`
    Duration    int    `yaml:"duration"`
}

func main() {
    if len(os.Args) != 3 {
        log.Fatalf("Usage: %s <input.yaml> <output.vtt>", os.Args[0])
    }

    input := os.Args[1]
    output := os.Args[2]

    f, err := os.ReadFile(input)
    if err != nil {
        log.Fatalf("read error: %v", err)
    }

    var slides []Slide
    if err := yaml.Unmarshal(f, &slides); err != nil {
        log.Fatalf("yaml error: %v", err)
    }

    out, err := os.Create(output)
    if err != nil {
        log.Fatalf("create error: %v", err)
    }
    defer out.Close()

    out.WriteString("WEBVTT\n\n")

    for i, slide := range slides {
        if slide.Voice == "" {
            continue
        }

        start := formatTime(slide.At)
        end := formatTime(slide.At + slide.Duration)

        fmt.Fprintf(out, "%d\n%s --> %s\n%s\n\n", i+1, start, end, slide.Voice)
    }

    log.Printf("✅ Wrote subtitles to %s", output)
}

func formatTime(ms int) string {
    d := time.Duration(ms) * time.Millisecond
    return fmt.Sprintf("%02d:%02d:%02d.%03d",
        int(d.Hours()), int(d.Minutes())%60, int(d.Seconds())%60, d.Milliseconds()%1000)
}

