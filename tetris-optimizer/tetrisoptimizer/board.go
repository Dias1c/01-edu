package tetrisoptimizer

type board struct {
	Width, Height int
	Field         [][]bool
}

func (gBoard *board) Constructor(width, height int) {
	field := make([][]bool, height)
	for i := 0; i < height; i++ {
		field[i] = make([]bool, width)
	}

	gBoard = &board{Width: width, Height: height, Field: field}
}

func (gBoard *board) Refactor() {
	field := make([][]bool, gBoard.Height)
	for i := 0; i < gBoard.Height; i++ {
		field[i] = make([]bool, gBoard.Width)
	}

	gBoard.Field = field
}
