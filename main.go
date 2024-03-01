package main

import "github.com/Prajvalbadiger/tic-tac-toe/internal/game"

func main() {
	board := game.NewBoard()
	board.GameLoop()
}
