package tetrisoptimizer

import (
	"fmt"
)

type board struct {
	Width int
	Field [][]byte
}

// Getting Min Board For Starting Match
func (gBoard *board) Constructor(tetrimonies []*tetrimony) {
	countTetrimonys := len(tetrimonies)
	wLength := getNearSqrtNum(countTetrimonys * 4)
	if wLength > 26 {
		_CloseProgram(_ErrorCountTetrimonys)
	}
	field := make([][]byte, wLength)
	for i := 0; i < wLength; i++ {
		field[i] = make([]byte, wLength)
	}
	gBoard.Width = wLength
	gBoard.Field = field
}

//Creating new Field With Current Width
func (gBoard *board) Refactor() {
	field := make([][]byte, gBoard.Width)
	for i := 0; i < gBoard.Width; i++ {
		field[i] = make([]byte, gBoard.Width)
	}
	gBoard.Field = field
}

//Just printing Board
func (bord *board) PrintBoard() {
	for _, line := range bord.Field {
		for _, v := range line {
			if v == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%s", string('@'+v))
			}
		}
		fmt.Println()
	}
}

//Check can you push tetrimony on board
//Yes = true, No = false
func (board *board) CanPushTetrimony(x, y int, figure *tetrimony) bool {
	for i := 0; i < figure.Height; i++ {
		posY := y + i
		for j := 0; j < figure.Width; j++ {
			posX := x + j
			if board.Field[posY][posX] != 0 && figure.Figure[i][j] {
				return false
			}
		}
	}
	return true
}

// can Panic OUT OF RANGE
func (board *board) PushTetrimony(x, y int, figure *tetrimony, step byte) {
	for i := 0; i < figure.Height; i++ {
		posY := y + i
		for j := 0; j < figure.Width; j++ {
			posX := x + j
			if figure.Figure[i][j] {
				if board.Field[posY][posX] == 0 {
					board.Field[posY][posX] += step
				} else {
					board.Field[posY][posX] += step
					board.PrintBoard()
					panic("Program try push on filled point")
				}
			}
		}
	}
}

// can Panic OUT OF RANGE
func (board *board) RemoveTetrimony(x, y int, figure *tetrimony, step byte) {
	for i := 0; i < figure.Height; i++ {
		posY := y + i
		for j := 0; j < figure.Width; j++ {
			posX := x + j
			if figure.Figure[i][j] {
				board.Field[posY][posX] -= step
			}
		}
	}
}
