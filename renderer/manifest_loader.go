package renderer

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"org-dasmlab/logutil"
)

// Logger instance from dasmlab standard
var log = logutil.InitLogger("manifest-loader")

// Slide represents a single slide in the tutorial
type Slide struct {
	Title       string          `yaml:"title"`
	Description string          `yaml:"description"`
	Voice       string          `yaml:"voice,omitempty"`
	Animations  []AnimationStep `yaml:"animations"`
}

// AnimationStep is a generic interface for all primitive steps
type AnimationStep struct {
	Type     string                 `yaml:"type"`
	Payload  map[string]interface{} `yaml:"payload"` // decoded as raw k-v map
	Timing   int                    `yaml:"at,omitempty"`
	Duration int                    `yaml:"duration,omitempty"`
}

// LoadManifest reads the YAML and returns parsed Slides
func LoadManifest(path string) ([]Slide, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorf("failed to read manifest: %v", err)
		return nil, err
	}
	var slides []Slide
	err = yaml.Unmarshal(data, &slides)
	if err != nil {
		log.Errorf("failed to parse YAML: %v", err)
		return nil, err
	}
	log.Infof("loaded manifest with %d slides", len(slides))
	return slides, nil
}

