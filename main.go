package main

import "github.com/gdamore/tcell"

func main() {
	view := NewView()

	quit := make(chan struct{})

	go func() {
		for {
			ev := view.screen.PollEvent()
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

	view.screen.Fini()
}
