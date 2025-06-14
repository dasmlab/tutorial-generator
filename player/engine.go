package player

import (
	"fmt"
	"time"

	"org-dasmlab/logutil"
	"org-dasmlab/renderer"
)

var log = logutil.InitLogger("player-engine")

// SimulatePlayback executes slide content sequentially
func SimulatePlayback(slides []renderer.Slide) {
	for i, slide := range slides {
		log.Infof("📽️ Slide %d: %s", i+1, slide.Title)
		fmt.Println(slide.Description)
		if slide.Voice != "" {
			log.Infof("🎤 Voice-over: %s", slide.Voice)
		}
		for _, step := range slide.Animations {
			log.Infof("▶️ Animate: %s at %ds", step.Type, step.Timing)
			time.Sleep(time.Duration(step.Timing) * time.Second)
			fmt.Printf("Animating: %s → %+v\n", step.Type, step.Payload)
		}
		fmt.Println("----")
	}
}

