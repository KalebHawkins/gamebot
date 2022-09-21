package gamebot

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
)

// WindowPidNotFoundError is returned when the window PID cannot be found.
type WindowPidNotFoundError struct {
	Applciation string
}

func (e *WindowPidNotFoundError) Error() string {
	return fmt.Sprintf("WindowPidNotFoundError: failed to find pid for %s", e.Applciation)
}

func (e *WindowPidNotFoundError) Is(tgt error) bool {
	_, ok := tgt.(*WindowPidNotFoundError)
	return ok
}

// NewWindowPidNotFound is returned if a window PID cannot be found.
func NewWindowPidNotFound(processName string) *WindowPidNotFoundError {
	return &WindowPidNotFoundError{
		Applciation: processName,
	}
}

// WindowPidGreaterThenOneError is returned when nPids > 1.
type WindowPidGreaterThenOneError struct {
	Applciation string
}

func (e *WindowPidGreaterThenOneError) Error() string {
	return fmt.Sprintf("WindowPidGreaterThenOneError: more then one PID was found for %s", e.Applciation)
}

func (e *WindowPidGreaterThenOneError) Is(tgt error) bool {
	_, ok := tgt.(*WindowPidGreaterThenOneError)
	return ok
}

// NewWindowPidGreaterThenOneError is returned when nPids > 1.
func NewWindowPidGreaterThenOneError(processName string) *WindowPidGreaterThenOneError {
	return &WindowPidGreaterThenOneError{
		Applciation: processName,
	}
}

type postiion struct {
	x, y int
}

type size struct {
	w, h int
}

type window struct {
	processName string
	hwnd        win.HWND
	title       string
	pid         int32
	position    postiion
	size        size
}

// This is how the window struct is to be displayed.
func (w *window) String() string {
	return fmt.Sprintf("(Pid: %d, X: %d, Y: %d, Width: %d, Height: %d)", w.pid, w.position.x, w.position.y, w.size.w, w.size.h)
}

// (w *Window) Title() returns the title of the window instance.
func (w *window) Title() string {
	return w.title
}

// (w *Window) Pid() returns the process id of the window instance.
func (w *window) Pid() int32 {
	return w.pid
}

// (w *Window) Position() returns the x and y coordinates of the upper-left corner of the window instance.
func (w *window) Position() (int, int) {
	return w.position.x, w.position.y
}

// (w *Window) Size() returns the width and height of the window instance.
func (w *window) Size() (int, int) {
	return w.size.w, w.size.h
}

// getWindow will search for a window by process name.
//
// If there are multiple pids found a WindowPidNotFound error.
//
// If there if more then one pid is found a WindowPidGreaterThenOneError is returned.
//
// It is possible for other errors to be returned.
func getWindow(procName string) (*window, error) {
	ids, err := robotgo.FindIds(procName)

	if err != nil {
		return nil, err
	}

	if len(ids) < 1 {
		return nil, NewWindowPidNotFound(procName)
	}

	if len(ids) > 1 {
		return nil, NewWindowPidGreaterThenOneError(procName)
	}

	id := ids[0]
	title := robotgo.GetTitle(id)
	x, y, w, h := robotgo.GetBounds(id)
	hwnd := robotgo.GetHWND()

	return &window{
		processName: procName,
		hwnd:        hwnd,
		title:       title,
		pid:         id,
		position: postiion{
			x: x,
			y: y,
		},
		size: size{
			w: w,
			h: h,
		},
	}, nil
}

// (w *window) IsActive returns true if the window instance is the active window otherwise false.
func (w *window) IsActive() bool {
	return robotgo.GetPID() == w.pid
}

// (w *window) SetActive set the window isnstance as the active window.
func (w *window) SetActive() {
	robotgo.ActivePID(w.pid)
}

// (w *window) Changed() check if the current window has changed size or position.
// If there is a probelm fetching the window the an error is returned.
func (w *window) Changed() (bool, error) {
	win, err := getWindow(w.processName)
	if err != nil {
		return false, err
	}

	if win.position.x != w.position.x || win.position.y != w.position.y {
		return true, err
	}

	if win.size.w != w.size.w || win.size.h != w.size.h {
		return true, err
	}

	return false, nil
}

// (b *Bot) Window() returns the bot's current window information.
func (b *Bot) Window() *window {
	b.config.botRWMut.RLock()
	defer b.config.botRWMut.RUnlock()

	return b.config.window
}

// (b *Bot) UpdateWindow sets the bot's window configuration in the event that the window changed size or position.
func (b *Bot) UpdateWindow() error {
	b.config.botRWMut.Lock()
	defer b.config.botRWMut.Unlock()

	windowChanged, err := b.config.window.Changed()
	if err != nil {
		return err
	}

	if windowChanged {
		win, err := getWindow(b.config.processName)
		if err != nil {
			return err
		}

		b.config.window = win
	}

	return nil
}
