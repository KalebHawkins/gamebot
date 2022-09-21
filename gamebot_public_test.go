package gamebot_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/KalebHawkins/gamebot"
)

func TestNewBot(t *testing.T) {
	proc := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(proc)

	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}

	posX, posY := b.Window().Position()
	sizeW, sizeH := b.Window().Size()

	want := 0
	if b.ProcessName() != proc {
		t.Errorf("expected process name %s, got %s", proc, b.ProcessName())
	}

	if posX != want || posY != want {
		t.Errorf("expected window positition %d, %d, got %d, %d", want, want, posX, posY)
	}

	if sizeW != want || sizeH != want {
		t.Errorf("expected window size %d, %d, got %d, %d", want, want, sizeW, sizeH)
	}
}
