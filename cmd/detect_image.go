package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/KalebHawkins/gamebot"
)

var procName *string
var imgPath *string

func parseFlags() {
	procName = flag.String("process", "", "the process name to target")
	imgPath = flag.String("image", "", "the image path to detect within the process window")
	flag.Parse()
}

func main() {
	parseFlags()

	// Create a new gamebot
	b, err := gamebot.NewBot(*procName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create bot: %s", err)
	}

	// Load an image to detect.
	targetImage, err := b.OpenImage(*imgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load image: %s", err)
	}

	go func() {
		// This function will open a window and continually update it
		// drawing a rectangle around the detected image. Press the 'q' key to quit the window.
		err = b.ShowDetectedImage("Debug Window", targetImage)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed during image detection: %s", err)
		}
	}()

	// Block until CTRL+C is pressed.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	<-sigs
}
