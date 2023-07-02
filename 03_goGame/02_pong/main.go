package main

import (
	"fmt"
	"os"

	// "github.com/gdamore/tcell/encoding"
	"github.com/gdamore/tcell"
)

// Plan
// capablity to render stuff on screen
// Draw paddles
// Playr movement
// Take care of paddle boundaries
// Draw ball
// Update ball movement
// Handle collisions
// Handle game over

func Print(s tcell.Screen, col, row int, width, height int, ch rune) {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			s.SetContent(col+c, row+r, ch, nil, tcell.StyleDefault)
		}
	}

}

func displayHelloWrod(screen tcell.Screen) {
	// w, h := screen.Size()
	screen.Clear()
	Print(screen, 0, 0, 8, 1, '*')
	Print(screen, 0, 0, 1, 8, '*')
	screen.Show()
}

func main() {
	// encoding.Register()
	screen := InitScreen()
	displayHelloWrod(screen)

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEnter {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}

func InitScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)
	return screen
}
