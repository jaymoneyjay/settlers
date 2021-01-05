package main

import (
	"github.com/gdamore/tcell"
)

// NewEngine instantiates a new game engine
func NewEngine() {
	engine := new(Engine)
	engine.Run()
}

// Run starts the game engine
func (e *Engine) Run() {
	quit := make(chan struct{})

	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				}
			}
		}
	}()

loop:
	for {
		select {
		case <-quit:
			break loop
		default:
			view.RefreshScreen()
		}
	}

	screen.Fini()
}
