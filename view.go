package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gdamore/tcell"
)

// NewView returns a pointer to a new instance of View
func NewView() *View {

	loader := NewLoader()
	boardTmpl := loader.load("boardAscii.txt")

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal("NewScreen error:", err)
	}

	err = screen.Init()
	if err != nil {
		log.Fatal("Screen Init error:", err)
	}

	screen.Clear()
	view := &View{
		screen: screen,
		boardTmpl: boardTmpl,
	}

	return view
}

func (view *View) initializeBoard(board *Board) {
	view.screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	styleBoard := tcell.StyleDefault.Foreground(tcell.ColorLightGray).Background(tcell.ColorBlack)

	// Draw board layout
	for y, line := range view.boardTmpl {
		for x, c := range line {
			view.screen.SetContent(x, y, c, nil, styleBoard)
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
				view.tcPrint(xPos, yPos, numberString, styleBoard)

				xPos -= 2
				yPos += 2
				resourceString := t.PrintResource()
				view.tcPrint(xPos, yPos, resourceString, styleBoard)
			}
		}
	}
	view.screen.Show()
}

func (view *View) drawBoard() {
	// Do nothing
}

// RefreshScreen updates the view of the screen
func (view *View) RefreshScreen() {
	view.drawBoard()
	time.Sleep(animationSpeed)
	view.screen.Show()
}

func (view *View) tcPrint(x, y int, text string, style tcell.Style) {
	for _, r := range text {
		view.screen.SetContent(x, y, r, nil, style)
		x++
	}
}
