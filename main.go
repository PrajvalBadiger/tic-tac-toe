package main

import (
	"fmt"

	"github.com/Prajvalbadiger/tic-tac-toe/internal/game"
)

func main() {

	fmt.Println("Welcome to Tic Tac Toe")
	board := game.NewBoard()
	board.GameLoop()
}
