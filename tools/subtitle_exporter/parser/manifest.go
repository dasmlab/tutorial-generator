package parser

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Slide defines the structure of a tutorial slide
type Slide struct {
	Title       string        `yaml:"title"`
	Description string        `yaml:"description"`
	Voice       string        `yaml:"voice"`
	At          float64       `yaml:"at"`
	Duration    float64       `yaml:"duration"`
	Animations  []interface{} `yaml:"animations"` // not needed for VTT, but parsed
}

// LoadSlides reads and parses slides from a YAML manifest
func LoadSlides(path string) ([]Slide, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var slides []Slide
	err = yaml.Unmarshal(file, &slides)
	if err != nil {
		return nil, err
	}

	return slides, nil
}

