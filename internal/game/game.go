package game

import (
	"fmt"
	"os"
)

func (pos *Position) isValid(b *Board) bool {
	return pos.X >= 0 && pos.X < Size && pos.Y >= 0 && pos.Y < Size && b.Curr[pos.X][pos.Y] == ' '
}

func (b *Board) GameLoop() {
	for {
		err := b.Play()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (b *Board) Play() error {
	var playing *Player
	// print board
	b.PrintBoard()

	// which player plays ?
	if b.Turn {
		// Player 2 playes
		playing = b.P2
	} else {
		// Player 1 playes
		playing = b.P1
	}
	fmt.Printf("Player %c plays\n", playing.Value)

	// get input
	pos := Position{}
	fmt.Scanf("%v %v", &pos.X, &pos.Y)

	if !pos.isValid(b) {
		return fmt.Errorf("Position is not valid (%d, %d)", pos.X, pos.Y)
	}

	b.Nmoves++
	// update board
	b.Curr[pos.X][pos.Y] = playing.Value
	b.Turn = !b.Turn

	// Check if the Game is over
	gameStatus := b.IsGameOver(&pos)

	switch gameStatus {
	case Draw:
		b.PrintBoard()
		fmt.Println("Draw")
		os.Exit(0)
	case Xwins:
		b.PrintBoard()
		fmt.Println("X wins")
		os.Exit(0)
	case Owins:
		b.PrintBoard()
		fmt.Println("O wins")
		os.Exit(0)
	case ContinueGame:
	}

	return nil
}

func (b *Board) IsGameOver(pos *Position) int {
	xSum := int('X' * Size)
	oSum := int('O' * Size)

	colSum := 0
	rowSum := 0
	diagSum := 0
	antiDiagSum := 0

	for i := 0; i < Size; i++ {
		colSum += b.Curr[pos.X][i]
		rowSum += b.Curr[i][pos.Y]
		diagSum += b.Curr[i][i]
		antiDiagSum += b.Curr[Size-i-1][i]
	}

	if rowSum == xSum || colSum == xSum || diagSum == xSum || antiDiagSum == xSum {
		return Xwins
	}

	if rowSum == oSum || colSum == oSum || diagSum == oSum || antiDiagSum == oSum {
		return Owins
	}

	// draw
	if b.Nmoves >= Size*Size {
		return Draw
	}

	// Game not yet over
	return ContinueGame
}
