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
	fmt.Println("Start Size:", wLength)
	gBoard.Width = wLength
	gBoard.Field = field
}

func (gBoard *board) Refactor() {
	field := make([][]byte, gBoard.Width)
	for i := 0; i < gBoard.Width; i++ {
		field[i] = make([]byte, gBoard.Width)
	}
	fmt.Println("Change Size:", gBoard.Width)
	gBoard.Field = field
}

var Try = 0

func (bord *board) PrintBoard() {
	fmt.Println("--- Print Board ---")
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

	// file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	// if err != nil {
	// 	_CloseProgram(err.Error())
	// }

	// writer := bufio.NewWriter(file)
	// _, err = writer.WriteString(fmt.Sprintf("--- Try: %v ---\n", Try))
	// if err != nil {
	// 	_CloseProgram(err.Error())
	// }
	// for _, line := range bord.Field {
	// 	for _, v := range line {
	// 		if v == 0 {
	// 			fmt.Fprint(writer, ".")
	// 		} else {
	// 			fmt.Fprintf(writer, "%s", string('@'+v))
	// 		}
	// 	}
	// 	fmt.Fprintln(writer)
	// }
	// writer.Flush()
	// Try++
	// file.Close()
}

// can Panic OUT OF RANGE
func (board *board) PushTetrimony(x, y int, figure *tetrimony, step byte) bool {
	result := true
	for i := 0; i < figure.Height; i++ {
		for j := 0; j < figure.Width; j++ {
			posX, posY := x+j, y+i
			if board.Field[posY][posX] == 0 {
				if figure.Figure[i][j] {
					board.Field[posY][posX] += step
				}
			} else {
				if figure.Figure[i][j] {
					board.Field[posY][posX] += step
					result = false
				}
			}
		}
	}
	return result
}

// can Panic OUT OF RANGE
func (board *board) RemoveTetrimony(x, y int, figure *tetrimony, step byte) bool {
	for i := 0; i < figure.Height; i++ {
		for j := 0; j < figure.Width; j++ {
			posX, posY := x+j, y+i
			if board.Field[posY][posX] == 0 {
				if figure.Figure[i][j] {
					board.Field[posY][posX] -= step
				}
			} else {
				if figure.Figure[i][j] {
					board.Field[posY][posX] -= step
				}
			}
		}
	}
	return true
}
