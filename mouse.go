package gamebot

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

type MouseButton string

const (
	Left       MouseButton = "left"
	Right      MouseButton = "center"
	Center     MouseButton = "right"
	WheelDown  MouseButton = "wheelDown"
	WheelUp    MouseButton = "wheelUp"
	WheelLeft  MouseButton = "wheelLeft"
	WheelRight MouseButton = "wheelRight"
)

// (b *Bot) MoveCursor simulates moving the cursor from it's current position to the x, y
// location on the screen. This simulates human-like movement. If you want to move x, y number
// of pixels from the current mouses position see `(b *Bot) MoveCursorRelative`.
// (x: 0, y: 0) represents the top left-hand corner of the screen.
func (b *Bot) MoveCursor(x, y int) {
	robotgo.MoveSmooth(x, y, 0.5, 1.0)
}

// (b *Bot) MoveCursorRelative simulates moving the cursor from it's current position
// by x and y number of pixels. This simulates human-like movement. x represents left and
// right movement while y represents up and down on the screen.
func (b *Bot) MoveCursorRelative(x, y int) {
	robotgo.MoveSmoothRelative(x, y, 0.5, 1.0)
}

// (b *Bot) SetCursor puts the cursor at the specified x, y position. This movement is nearly instant
// and does not simulate human-like movement.
func (b *Bot) SetCursor(x, y int) {
	robotgo.Move(x, y)
}

// (b *Bot) MoveClick puts the cursor at the specified x, y position then clicks the specified mouse button. This movement is nearly instant
// and does not simulate human-like movement. Reference https://github.com/go-vgo/robotgo/blob/master/docs/keys.md for keycodes.
func (b *Bot) MoveCursorClick(x, y int, btn MouseButton, doubleClick bool) {
	robotgo.MoveClick(x, y, btn, doubleClick)
}

// (b *Bot) MoveCursorSmoothClick puts the cursor at the specified x, y position then clicks the specified mouse button.
// This movement simulates human-like movement. Reference https://github.com/go-vgo/robotgo/blob/master/docs/keys.md for keycodes.
func (b *Bot) MoveCursorSmoothClick(x, y int, btn MouseButton, doubleClick bool) {
	robotgo.MovesClick(x, y, btn, doubleClick)
}

// (b *Bot) Click click the specified mouse button at the current location of the cursor.
// Reference https://github.com/go-vgo/robotgo/blob/master/docs/keys.md for keycodes.
func (b *Bot) Click(btn MouseButton, doubleClick bool) {
	robotgo.Click(btn, doubleClick)
}

// (b *Bot) MousePosition returns the mouse's current x, y coordinates.
func (b *Bot) MousePosition() (int, int) {
	return robotgo.GetMousePos()
}

// (b *Bot) GetPixelColor return the color of the pixel at the x, y coordinates of the screen.
func (b *Bot) GetPixelColor(x, y int) string {
	return robotgo.GetPixelColor(x, y)
}

// (b *Bot) MousePress puts the specified mouse button in a down state. To release the button use `(b *Bot) MousePress`.
func (b *Bot) MousePress(btn MouseButton) {
	b.config.botRWMut.RLock()
	defer b.config.botRWMut.RUnlock()

	robotgo.Toggle(string(btn))

	mouseButton := fmt.Sprintf("mouse%s", btn)
	b.config.keysDown[mouseButton] = true
}

// (b *Bot) MouseRelease puts the specified mouse button in an up state.
func (b *Bot) MouseRelease(btn MouseButton) {
	b.config.botRWMut.RLock()
	defer b.config.botRWMut.RUnlock()

	robotgo.Toggle(string(btn))

	mouseButton := fmt.Sprintf("mouse%s", btn)
	delete(b.config.keysDown, mouseButton)
}
