package main

// BoardCreator is responsible for seeding different boards
type BoardCreator struct {
}

// NewBoardCreator returns a pointer to a new instance of BoardCreator
func NewBoardCreator() *BoardCreator {
	creator := &BoardCreator{}
	return creator
}

// SeedDefault initializes the specified board with a default board setup
func (s *BoardCreator) SeedDefault(board *Board) {

	// Initialize resources
	board.tiles[3][1] = NewTile(wood, 8)
	board.tiles[1][2] = NewTile(grain, 3)
	board.tiles[2][2] = NewTile(brick, 11)
	board.tiles[3][2] = NewTile(wood, 9)
	board.tiles[4][2] = NewTile(brick, 4)
	board.tiles[5][2] = NewTile(ore, 9)
	board.tiles[1][3] = NewTile(wood, 10)
	board.tiles[2][3] = NewTile(wool, 8)
	board.tiles[3][3] = NewTile(brick, 6)
	board.tiles[4][3] = NewTile(desert, 0)
	board.tiles[5][3] = NewTile(wool, 10)
	board.tiles[1][4] = NewTile(grain, 6)
	board.tiles[2][4] = NewTile(grain, 12)
	board.tiles[3][4] = NewTile(wool, 4)
	board.tiles[4][4] = NewTile(ore, 2)
	board.tiles[5][4] = NewTile(brick, 12)
	board.tiles[2][5] = NewTile(wood, 11)
	board.tiles[3][5] = NewTile(wool, 5)
	board.tiles[4][5] = NewTile(grain, 5)

	// Initialize water
	board.tiles[3][0] = NewTile(water, 0)
	board.tiles[1][1] = NewTile(water, 0)
	board.tiles[2][1] = NewTile(water, 0)
	board.tiles[4][1] = NewTile(water, 0)
	board.tiles[5][1] = NewTile(water, 0)
	board.tiles[0][2] = NewTile(water, 0)
	board.tiles[6][2] = NewTile(water, 0)
	board.tiles[0][3] = NewTile(water, 0)
	board.tiles[6][3] = NewTile(water, 0)
	board.tiles[0][4] = NewTile(water, 0)
	board.tiles[6][4] = NewTile(water, 0)
	board.tiles[0][5] = NewTile(water, 0)
	board.tiles[1][5] = NewTile(water, 0)
	board.tiles[5][5] = NewTile(water, 0)
	board.tiles[6][5] = NewTile(water, 0)
	board.tiles[2][6] = NewTile(water, 0)
	board.tiles[3][6] = NewTile(water, 0)
	board.tiles[4][6] = NewTile(water, 0)
}
