package gamebot

import "github.com/go-vgo/robotgo"

type KeyState string

const (
	Down KeyState = "down"
	Up   KeyState = "up"
)

// (b *Bot) PressKey toggles a key on the keyboard. This will put the key in a down state until ReleaseKey is called.
// Reference https://github.com/go-vgo/robotgo/blob/master/docs/keys.md for full list of keycodes.
func (b *Bot) PressKey(key string) {
	b.config.botRWMut.Lock()
	defer b.config.botRWMut.Unlock()

	robotgo.KeyToggle(key)
	b.config.keysDown[key] = true
}

// (b *Bot) ReleaseKey toggles a key on the keyboard. This will put the key in a down state until ReleaseKey is called.
// Reference https://github.com/go-vgo/robotgo/blob/master/docs/keys.md for full list of keycodes.
func (b *Bot) ReleaseKey(key string) {
	b.config.botRWMut.Lock()
	defer b.config.botRWMut.Unlock()

	robotgo.KeyToggle(key, string(Up))
	delete(b.config.keysDown, key)
}

// (b *Bot) IsKeyDown returns true is the specified key is in a down state.
//
// Note that mouse keys are prefixed with the string mouse e.g. mouseleft, mouseright, mousecenter
func (b *Bot) IsKeyDown(key string) bool {
	b.config.botRWMut.RLock()
	defer b.config.botRWMut.RUnlock()

	_, ok := b.config.keysDown[key]
	return ok
}

// (b *Bot) KeyTap will press and release a key.
// The `args` parameter represents special characters that may need to be pressed alongside the primary key, e.g shift.
// Reference https://github.com/go-vgo/robotgo/blob/master/docs/keys.md for full list of keycodes.
func (b *Bot) KeyTap(key string) {
	robotgo.KeyToggle(key)
	b.MilliSleep(300)
	robotgo.KeyToggle(key, string(Up))
}

// (b *Bot) KeysDown returns a slice of keys currently in the down position.
// If there are no keys in a down state this function returns an empty slice.
func (b *Bot) KeysDown() []string {
	b.config.botRWMut.Lock()
	defer b.config.botRWMut.Unlock()

	keys := make([]string, 0)
	for k := range b.config.keysDown {
		keys = append(keys, k)
	}

	return keys
}
