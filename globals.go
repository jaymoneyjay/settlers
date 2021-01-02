package main

import (
	"time"

	"github.com/gdamore/tcell"
)

const (
	animationSpeed = 10 * time.Millisecond
)

type (
	// View is the display engine
	View struct {
		boardTmpl []string
		screen    tcell.Screen
	}

	// TxtLoader loads and stores content of txt files line by line
	TxtLoader struct {
	}
)
