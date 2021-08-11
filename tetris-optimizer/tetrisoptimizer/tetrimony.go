package tetrisoptimizer

import (
	"bufio"
	"fmt"
	"os"
)

type tetrimony struct {
	Figure        [][]bool
	Width, Height int
}

func (tetrim *tetrimony) Constructor(lines []string) {
	matrix := make([][]rune, 4)
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	figure := checkAndGetMatrixForFigure(matrix)
	tetrim.Figure = figure
	tetrim.Height = len(figure)
	tetrim.Width = len(figure[0])
}

func _GetTetrimonysFromFile(filepath string) []*tetrimony {
	file, err := os.Open(filepath)
	if err != nil {
		_CloseProgram(err.Error())
	}
	tetrimonies := []*tetrimony{}
	reader := bufio.NewScanner(file)
	skip, take := 5, 1
	lines := []string{}
	for reader.Scan() {
		line := reader.Text()
		if take != skip {
			if len(line) != 4 {
				_CloseProgram(_ErrorIncorrectTetrimonyInFile)
			}
			lines = append(lines, line)
			take++
		} else {
			if len(line) != 0 {
				_CloseProgram(_ErrorIncorrectTetrimonyInFile)
			}
			take = 1
			tetrimony := tetrimony{}
			tetrimony.Constructor(lines)
			tetrimonies = append(tetrimonies, &tetrimony)

			lines = make([]string, 0)
		}
	}
	if len(lines) == 4 {
		tetrimony := tetrimony{}
		tetrimony.Constructor(lines)
		tetrimonies = append(tetrimonies, &tetrimony)
	} else {
		_CloseProgram(_ErrorIncorrectTetrimonyInFile)
	}
	return tetrimonies
}

func getCornerPositionsMatrix(matrix [][]rune) (int, int, int, int) {
	left, top, right, bottom := -1, -1, -1, -1
	isFirst := true
	cnt := 0
	for y, line := range matrix {
		for x, symbol := range line {
			if symbol == '.' || symbol == '#' {
				if symbol == '#' {
					if isFirst {
						left, right = x, x
						top, bottom = y, y
						isFirst = false
					} else {
						left = getMin(left, x)
						right = getMax(right, x)
						top = getMin(top, y)
						bottom = getMax(bottom, y)
					}
					cnt++
				}
			} else {
				msg := fmt.Sprintf(_ErrorIncorrectSymbolInFile, string(symbol))
				_CloseProgram(msg)
			}
		}
	}
	if cnt != 4 {
		_CloseProgram(_ErrorIncorrectTetrimonyInFile)
	}
	return left, top, right, bottom
}

//Проверяет матрицу меняя в нем данные
func checkMatrix(matrix [][]rune) {
	square, cntIslands := 0, 0
	for y, line := range matrix {
		for x, symbol := range line {
			if symbol == '#' {
				cntIslands++
				square = getSquareFromIsland(matrix, 4, 4, y, x)
			}
		}
	}
	if cntIslands != 1 || square != 4 {
		_CloseProgram(_ErrorIncorrectTetrimonyInFile)
	}
}

func checkAndGetMatrixForFigure(matrix [][]rune) [][]bool {
	left, top, right, bottom := getCornerPositionsMatrix(matrix)

	width, height := right-left+1, bottom-top+1
	// fmt.Printf("Width: %v,Height: %v\n", width, height)
	// fmt.Printf("L: %v, T: %v, R: %v, B: %v\n", left, top, right, bottom)
	result := make([][]bool, height)
	for y, i := top, 0; y <= bottom; y++ {
		result[i] = make([]bool, width)
		for x, j := left, 0; x <= right; x++ {
			if matrix[y][x] == '#' {
				result[i][j] = true
			}
			j++
		}
		i++
	}
	checkMatrix(matrix)
	return result
}

func getSquareFromIsland(matrix [][]rune, width, height, y, x int) int {
	if y < 0 || y == height || x < 0 || x == width {
		return 0
	}
	cnt := 0
	if matrix[y][x] == '#' {
		matrix[y][x] = '.'
		cnt++
		cnt += getSquareFromIsland(matrix, width, height, y+1, x)
		cnt += getSquareFromIsland(matrix, width, height, y-1, x)
		cnt += getSquareFromIsland(matrix, width, height, y, x+1)
		cnt += getSquareFromIsland(matrix, width, height, y, x-1)
	}
	return cnt
}
