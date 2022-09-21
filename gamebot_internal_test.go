package gamebot

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWindowFuncs(t *testing.T) {
	proc := filepath.Base(os.Args[0])
	win, err := getWindow(proc)

	t.Run("Test getWindow", func(t *testing.T) {
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}

		if win == nil {
			t.Errorf("expected non nil window, got %v", win)
		}
	})

	t.Run("Test window.IsActive", func(t *testing.T) {
		want := false

		if v := win.IsActive(); v != want {
			t.Errorf("expected %t, got %t", want, v)
		}
	})

	t.Run("Test window.Changed", func(t *testing.T) {
		want := false
		if v, _ := win.Changed(); v != want {
			t.Errorf("expected %t, got %t", want, v)
		}

		saveState := win
		defer func() { win = saveState }()

		testChange := func(want bool, msg string) {
			if v, _ := win.Changed(); v != want {
				t.Errorf("expected %t, got %t: %s", want, v, msg)
			}
		}

		want = true
		win.size.w = 1
		testChange(want, "changed width")
		win.size.h = 1
		testChange(want, "changed height")
		win.position.x = 1
		testChange(want, "changed position x")
		win.position.y = 1
		testChange(want, "changed position y")
	})
}

func TestBotFuncs(t *testing.T) {
	proc := filepath.Base(os.Args[0])
	b, err := NewBot(proc)

	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	t.Run("Test bot.ProcessName", func(t *testing.T) {
		if v := b.ProcessName(); v != b.config.processName {
			t.Errorf("expected %s, got %s", proc, v)
		}
	})

	t.Run("Test bot.UpdateWindow", func(t *testing.T) {
		saveState := b.config.window
		defer func() { b.config.window = saveState }()

		b.config.window.position.x = 1
		b.config.window.position.x = 2
		err := b.UpdateWindow()
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}

		if b.config.window.position.x != 0 && b.config.window.position.y != 0 {
			t.Errorf("expected window position of %d, %d, got %d, %d", 0, 0, b.config.window.position.x, b.config.window.position.y)
		}

		b.config.window.size.w = 1
		b.config.window.size.h = 2
		b.UpdateWindow()
		if b.config.window.size.w != 0 && b.config.window.size.h != 0 {
			t.Errorf("expected window size of %d, %d, got %d, %d", 0, 0, b.config.window.size.w, b.config.window.size.h)
		}
	})

}
