package game

import (
	"fmt"
)

const Size int = 3
const (
	Draw         int = 0
	Xwins        int = 1
	Owins        int = 2
	ContinueGame int = 3
)

type Player struct {
	Value int
}

type Board struct {
	Curr   [Size][Size]int
	P1     *Player
	P2     *Player
	Nmoves int
	Turn   bool
}

type Position struct {
	X int
	Y int
}

func NewBoard() *Board {
	b := &Board{
		P1: &Player{Value: 'X'},
		P2: &Player{Value: 'O'},
	}

	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			b.Curr[i][j] = ' '
		}
	}

	return b
}

// TODO: printing the board doesn't depend on Size, this method only prints for Sixe = 3
func (b *Board) PrintBoard() {
	fmt.Printf("%+v\n", b)
	fmt.Printf("%c|%c|%c\r\n", b.Curr[0][0], b.Curr[0][1], b.Curr[0][2])
	fmt.Println("-+-+-\r")
	fmt.Printf("%c|%c|%c\r\n", b.Curr[1][0], b.Curr[1][1], b.Curr[1][2])
	fmt.Println("-+-+-")
	fmt.Printf("%c|%c|%c\r\n", b.Curr[2][0], b.Curr[2][1], b.Curr[2][2])
}
