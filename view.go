package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell"
)

// NewView returns a pointer to a new instance of View
func NewView() *View {
	view := new(View)
	loader := NewLoader()
	loader.load("boardAscii.txt")
	view.boardTmpl = loader.lines

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal("NewScreen error:", err)
	}

	view.screen = screen
	view.screen.Init()
	view.screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	view.screen.Clear()

	return view
}

func (v *View) drawBoard() {
	styleBoard := tcell.StyleDefault.Foreground(tcell.ColorLightGray).Background(tcell.ColorBlack)
	for y, line := range v.boardTmpl {
		for x, c := range line {
			v.screen.SetContent(x, y, c, nil, styleBoard)
		}
	}
}

// RefreshScreen updates the view of the screen
func (v *View) RefreshScreen() {
	v.drawBoard()
	time.Sleep(animationSpeed)
	v.screen.Show()
}
