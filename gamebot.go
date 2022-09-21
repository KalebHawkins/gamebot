package gamebot

import (
	"math/rand"
	"sync"

	"github.com/go-vgo/robotgo"
	"gocv.io/x/gocv"
)

const (
	// screenCaptureDelayMs represents how long robotgo should sleep, in milliseconds, after capturing the screen before scanning the screen for pixels.
	screenCaptureDelayMs = 300

	// cvMatchMode represents the default value for gamebot's opencv template matching algorithm.
	cvMatchMode = gocv.TmCcoeffNormed
)

type Bot struct {
	config *botConfig
}

type botConfig struct {
	botRWMut sync.RWMutex

	processName string
	window      *window

	screenCaptureDelayMs int

	// keysDown keeps track of all the keys in a down state.
	keysDown map[string]bool

	cvMatchMode gocv.TemplateMatchMode
}

// NewBot create a new bot instance.
func NewBot(processName string) (*Bot, error) {
	win, err := getWindow(processName)
	if err != nil {
		return nil, err
	}

	config := &botConfig{}
	config.processName = processName
	config.window = win
	config.screenCaptureDelayMs = screenCaptureDelayMs
	config.cvMatchMode = cvMatchMode

	// keysDown is a map of strings that are currently in the 'down' or 'pressed' state.
	// Mouse keys are prefixed with the string `mouse` to be able to distinguish between keyboard's left and right keys
	// vs the mouses left and right buttons.
	config.keysDown = make(map[string]bool)

	b := &Bot{}
	b.config = config
	return b, nil
}

// (b *Bot) ProcessName() return the bot's currently configured processName.
func (b *Bot) ProcessName() string {
	b.config.botRWMut.RLock()
	defer b.config.botRWMut.RUnlock()

	return b.config.processName
}

// (b *Bot) MilliSleep will pause program operation for the specified number of milliseconds.
func (b *Bot) MilliSleep(ms int) {
	robotgo.MilliSleep(ms)
}

// (b *Bot) Sleep will pause program operation for the specified number of seconds.
func (b *Bot) Sleep(s int) {
	robotgo.Sleep(s)
}

// (b *Bot) RandomInt generate random integers within a range of min and max values (inclusive).
func (b *Bot) RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
