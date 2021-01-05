package main

import (
	"log"
)

// NewTile returns a pointer to a new instantiation of Tile
func NewTile(res Resource, num int) *Tile {
	t := &Tile{
		resource: res,
		number:   num,
	}
	return t
}

// PrintResource returns the resource of the tile in string representation
func (t *Tile) PrintResource() string {
	switch t.resource {
	case wood:
		return "wood"
	case ore:
		return "ore"
	case wool:
		return "wool"
	case grain:
		return "grain"
	case brick:
		return "brick"
	case desert:
		return "desert"
	case water:
		return "water"
	default:
		log.Printf("This resource does not exist: %d", t.resource)
		return "Error"
	}
}
