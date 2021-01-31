package main

// NewBoard creates a new instantiation of Board
func NewBoard() *Board {

	// Initialize tiles
	t := make([][]*Tile, 7, 7)
	for i := range t {
		t[i] = make([]*Tile, 7, 7)
	}

	// INitialize edges
	e := make([][]*Edge, 7, 7)
	for i := range e {
		e[i] = make([]*Edge, 16, 16)
	}

	// Initialize vertices
	v := make([][]*Vertex, 16, 16)
	for i := range v {
		v[i] = make([]*Vertex, 16, 16)
	}

	return &Board{
		tiles:    t,
		edges:    e,
		vertices: v,
	}
}

/*
// InitializeBoard sets up a board with the specified configuration
func InitializeBoard(config string) {

}


// GetVertices returns the adjactent vertices to the specified tile
func GetVertices(tile Tile) []Vertex {
	return {}
}

// GetEdges returns the adjactent edges to the specified tile
func GetEdges(tile Tile) []Edge {
	return {}
}
*/
