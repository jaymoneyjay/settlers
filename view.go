package main

import (
	"fmt"
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

	view.initializeBoard()
}

func (view *View) initializeBoard() {
	screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	styleBoard := tcell.StyleDefault.Foreground(tcell.ColorLightGray).Background(tcell.ColorBlack)

	// Draw board layout
	for y, line := range view.boardTmpl {
		for x, c := range line {
			screen.SetContent(x, y, c, nil, styleBoard)
		}
	}

	// Draw board content
	for x := range board.tiles {
		for y, t := range board.tiles[x] {
			if t != nil && t.resource != water && t.resource != desert {
				xPos := x*9 + 6
				yPos := y*6 + 2
				if x%2 == 0 {
					yPos -= 3
				}
				numberString := fmt.Sprintf("%d", t.number)
				tcPrint(xPos, yPos, numberString, styleBoard)

				xPos -= 2
				yPos += 2
				resourceString := t.PrintResource()
				tcPrint(xPos, yPos, resourceString, styleBoard)
			}
		}
	}
	screen.Show()
}

func (view *View) drawBoard() {
	// Do nothing
}

// RefreshScreen updates the view of the screen
func (view *View) RefreshScreen() {
	view.drawBoard()
	time.Sleep(animationSpeed)
	screen.Show()
}

func tcPrint(x, y int, text string, style tcell.Style) {
	for _, r := range text {
		screen.SetContent(x, y, r, nil, style)
		x++
	}
}
