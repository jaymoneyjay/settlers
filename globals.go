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
	}

	// TxtLoader loads and stores content of txt files line by line
	TxtLoader struct {
	}

	// Engine is the game engine of the settlers of catan game
	Engine struct {
	}
)

var (
	engine *Engine
	view   *View
	loader *TxtLoader
	screen tcell.Screen
)
