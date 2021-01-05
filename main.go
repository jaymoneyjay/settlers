package main

func main() {
	NewBoard()
	NewBoardCreator().SeedDefault(board)
	NewView()
	NewEngine()
}
