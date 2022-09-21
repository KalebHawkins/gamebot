package gamebot_test

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/KalebHawkins/gamebot"
)

func ExampleNewBot() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// return the process name the bot is attached to.
	fmt.Println(b.ProcessName())
	// output: gamebot.test.exe
}

func ExampleBot_Sleep() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// Bot will pause operation for 1 seconds
	b.Sleep(1)
}

func ExampleBot_MilliSleep() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// Bot will pause operation for 100 milliseconds
	b.MilliSleep(100)
}

func ExampleBot_RandomInt() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	rand.Seed(0)

	// get a random integer between 5 and 10 (he min and max values are inclusive.
	rn := b.RandomInt(5, 10)
	fmt.Println(rn)
	// output: 5
}

func ExampleBot_PressKey() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// Press the 'w' key to a down state
	b.PressKey("w")
	// Release the 'w' key
	b.ReleaseKey("w")
}

func ExampleBot_ReleaseKey() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// Press the 'w' key to a down state
	b.PressKey("w")
	// Release the 'w' key
	b.ReleaseKey("w")
}

func ExampleBot_IsKeyDown() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// Press the 'w' key to a down state
	b.PressKey("w")

	// Check if a key is pushed down.
	if b.IsKeyDown("w") {
		fmt.Println("w key is down")
	}

	b.ReleaseKey("w")
	//output: w key is down
}

func ExampleBot_KeysDown() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	b.PressKey("w")
	b.PressKey("a")
	b.PressKey("s")
	b.PressKey("d")

	// Get a slice of keys in down position
	depressedKeys := b.KeysDown()

	for _, v := range depressedKeys {
		fmt.Println(v)
	}

	b.ReleaseKey("w")
	//output: w
	// a
	// s
	// d
}

func ExampleBot_MoveCursor() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// Drag the cursor position to the top-left hand corner
	b.MoveCursor(0, 0)
}

func ExampleBot_MoveCursorRelative() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// Drag the cursor position from the current position
	// 10 pixels left and 10 pixels down.
	b.MoveCursor(-10, 10)
}

func ExampleBot_SetCursor() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// Instantly set the location of the cursor to 100, 100
	b.SetCursor(100, 100)
}

func ExampleBot_MoveCursorClick() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// Instantly set the location of the cursor to 100, 100
	// and click the left mouse button once.
	b.MoveCursorClick(100, 100, gamebot.Left, false)

	// Instantly set the location of the cursor to 100, 100
	// and double click the left mouse
	b.MoveCursorClick(100, 100, gamebot.Left, true)
}

func ExampleBot_MoveCursorSmoothClick() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// Drag the location of the cursor to 100, 100
	// and click the left mouse button once.
	b.MoveCursorClick(100, 100, gamebot.Left, false)

	// Drag the location of the cursor to 100, 100
	// and double click the left mouse
	b.MoveCursorClick(100, 100, gamebot.Left, true)
}

func ExampleBot_Click() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// Single click the right mouse button
	b.Click(gamebot.Right, false)
}

func ExampleBot_MousePosition() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	mpx, mpy := b.MousePosition()
	fmt.Println(mpx, mpy)
	// ouptut: 0 0
}

func ExampleBot_MousePress() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	b.MousePress(gamebot.Left)

	if b.IsKeyDown("mouseleft") {
		fmt.Println("left mouse key is down")
	}

	b.MouseRelease(gamebot.Left)
	//output: left mouse key is down
}

func ExampleBot_MouseRelease() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	b.MousePress(gamebot.Left)

	if b.IsKeyDown("mouseleft") {
		fmt.Println("left mouse key is down")
	}

	b.MouseRelease(gamebot.Left)
	//output: left mouse key is down
}

func Examplewindow_Title() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// return the window title
	fmt.Println(b.Window().Title())
}

func Examplewindow_Pid() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// return the process id
	fmt.Println(b.Window().Pid())
}

func Examplewindow_Position() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// get the top-left corner coordinates of the window.
	px, py := b.Window().Position()
	fmt.Println(px, py)
}

func Examplewindow_Size() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	// get the top-left corner coordinates of the window.
	width, height := b.Window().Position()
	fmt.Println(width, height)
}

func Examplewindow_IsActive() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	if !b.Window().IsActive() {
		b.Window().SetActive()
	}
}

func Examplewindow_SetActive() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	if !b.Window().IsActive() {
		b.Window().SetActive()
	}
}

func Examplewindow_Changed() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	didChange, err := b.Window().Changed()
	if err != nil {
		panic(err)
	}

	if didChange {
		b.UpdateWindow()
	}
}

func ExampleBot_UpdateWindow() {
	procName := filepath.Base(os.Args[0])
	b, err := gamebot.NewBot(procName)

	if err != nil {
		panic(err)
	}

	didChange, err := b.Window().Changed()
	if err != nil {
		panic(err)
	}

	if didChange {
		b.UpdateWindow()
	}
}
