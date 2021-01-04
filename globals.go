package main

import (
	"time"

	"github.com/gdamore/tcell"
)

const (
	animationSpeed = 10 * time.Millisecond
)

// Resource is a resource players can yield
// Example: wood, ore etc.
type Resource int

const (
	desert Resource = iota - 1
	water
	wood
	brick
	lumber
	wool
	grain
	ore
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

	// Tile is a single hexagon of the game board representing a resource
	Tile struct {
		resource Resource
		number   int
	}

	// Vertex is a node in the middle of three tiles
	// represents the location for buildings
	Vertex struct {
	}

	// Edge is a link between two vertices
	// represents the location to build streets
	Edge struct {
	}

	// Board models the settlers of catan game boar
	Board struct {
		tiles    [][]Tile
		edges    [][]Edge
		vertices [][]Vertex
	}
)

var (
	engine *Engine
	view   *View
	loader *TxtLoader
	screen tcell.Screen
	board  *Board
)
