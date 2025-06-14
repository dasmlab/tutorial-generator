package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Slide struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Voice       string `yaml:"voice"`
}

func main() {
	log := logrus.New()
	if len(os.Args) != 2 {
		log.Fatal("Usage: subtitle_writer <file.yaml>")
	}

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to read YAML: %v", err)
	}

	var slides []Slide
	err = yaml.Unmarshal(content, &slides)
	if err != nil {
		log.Fatalf("YAML parsing error: %v", err)
	}

	log.Infof("📝 Writing subtitles to 'subtitles.vtt'")

	out, _ := os.Create("subtitles.vtt")
	fmt.Fprintln(out, "WEBVTT\n")

	start := 0
	for i, slide := range slides {
		if slide.Voice != "" {
			end := start + 4
			fmt.Fprintf(out, "%d\n00:00:%02d.000 --> 00:00:%02d.000\n%s\n\n",
				i+1, start, end, slide.Voice)
			start = end + 1
		}
	}
	log.Infof("✅ Done.")
}

