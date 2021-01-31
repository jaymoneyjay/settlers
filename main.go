package main

func main() {
	board := NewBoard()
	board.SeedDefault()

	view := NewView()
	view.initializeBoard(board)

	engine := NewEngine(view)
	engine.Run()
}
