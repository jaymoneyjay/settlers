package main

import (
	"github.com/gdamore/tcell"
)

// NewEngine instantiates a new game engine
func NewEngine(view *View) *Engine {
	engine := &Engine{
		view: view,
	}
	return engine
}

// Run starts the game engine
func (e *Engine) Run() {
	quit := make(chan struct{})

	go func() {
		for {
			ev := e.view.screen.PollEvent()
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
			e.view.RefreshScreen()
		}
	}

	e.view.screen.Fini()
}
