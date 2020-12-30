package main

import "github.com/gdamore/tcell"

const ()

type (
	// View is the display engine
	View struct {
		boardTmpl []string
		screen    tcell.Screen
	}

	// TxtLoader loads and stores content of txt files line by line
	TxtLoader struct {
		lines []string
	}
)
