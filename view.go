package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell"
)

// NewView returns a pointer to a new instance of View
func NewView() {

	// Initialize globals
	view = new(View)
	loader = NewLoader()

	var err error

	view.boardTmpl = loader.load("boardAscii.txt")

	screen, err = tcell.NewScreen()
	if err != nil {
		log.Fatal("NewScreen error:", err)
	}

	err = screen.Init()
	if err != nil {
		log.Fatal("Screen Init error:", err)
	}

	screen.Clear()
}

func (view *View) drawBoard() {
	screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	styleBoard := tcell.StyleDefault.Foreground(tcell.ColorLightGray).Background(tcell.ColorBlack)
	for y, line := range view.boardTmpl {
		for x, c := range line {
			screen.SetContent(x, y, c, nil, styleBoard)
		}
	}
}

// RefreshScreen updates the view of the screen
func (view *View) RefreshScreen() {
	view.drawBoard()
	time.Sleep(animationSpeed)
	screen.Show()
}
