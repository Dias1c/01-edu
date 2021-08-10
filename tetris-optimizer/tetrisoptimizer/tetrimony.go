package tetrisoptimizer

import (
	"bufio"
	"os"
)

type tetrimony struct {
	Figure [][]bool
}

func _GetTetrimonysFromFile(filepath string) []tetrimony {
	file, err := os.Open(filepath)
	if err != nil {
		_CloseProgram(err.Error())
	}
	tetrimonies := []tetrimony{}
	reader := bufio.NewScanner(file)
	skip, take := 5, 1
	lines := []string{}
	for reader.Scan() {
		line := reader.Text()
		if take != skip {
			lines = append(lines, line)
			take++
		} else {
			if len(line) != 0 {
				_CloseProgram("Incorrect tetrimonies on file!")
			}
			take = 1
			tetrimony := tetrimony{}
			tetrimony.Constructor(lines)
			tetrimonies = append(tetrimonies, tetrimony)

			lines = make([]string, 0)
		}
	}
	if len(lines) == 4 {
		tetrimony := tetrimony{}
		tetrimony.Constructor(lines)
		tetrimonies = append(tetrimonies, tetrimony)
	}
	return tetrimonies
}

func (tetrim *tetrimony) Constructor(lines []string) {

}

func (tetrim *tetrimony) CheckTetrimony() {

}
