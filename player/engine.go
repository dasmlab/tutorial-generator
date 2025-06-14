package player

import (
	"encoding/json"
	"os"
	"sort"
	"tutorial-generator/renderer"
	"tutorial-generator/logutil"
)

var log = logutil.InitLogger("player-engine")

// TimelineCommand represents an actionable step in the frontend
type TimelineCommand struct {
	At       int64                  `json:"at"`
	Duration int64                  `json:"duration"`
	Type     string                 `json:"type"`
	Params   map[string]interface{} `json:"params"`
}

// BuildTimeline takes parsed slides and flattens them into time-sequenced commands
func BuildTimeline(slides []renderer.Slide) []TimelineCommand {
	var timeline []TimelineCommand

	for _, slide := range slides {
		for _, step := range slide.Animations {
			timeline = append(timeline, TimelineCommand{
				At:       step.At,
				Duration: step.Duration,
				Type:     step.Type,
				Params:   step.Params,
			})
		}
	}

	// Sort by 'At' time
	sort.Slice(timeline, func(i, j int) bool {
		return timeline[i].At < timeline[j].At
	})

	log.Infof("Built timeline with %d steps", len(timeline))
	return timeline
}

// DumpTimelineToFile writes timeline to JSON for front-end
func DumpTimelineToFile(timeline []TimelineCommand, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(timeline)
}

