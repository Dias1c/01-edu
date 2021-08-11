package tetrisoptimizer

import (
	"fmt"
	"os"
)

func Program(args []string) {
	_CheckForCountArgs(args)
	filepath := args[0]
	// fmt.Printf("FilePath: %v\n", filepath)
	tetrimonies := _GetTetrimonysFromFile(filepath)
	result := Match(tetrimonies)
	result.PrintBoard()
}

func Match(tetrimonies []*tetrimony) *board {
	result := &board{}
	result.Constructor(tetrimonies)
	isMatch := false
	for !isMatch {
		isMatch = _FindSmallestSquare(tetrimonies, result, 1)
		if !isMatch {
			result.Width++
			result.Refactor()
		}
	}
	return result
}

func _CheckForCountArgs(args []string) {
	if len(args) != 1 {
		_CloseProgram(_ErrorCountArgsNE1)
	}
}

func _CloseProgram(msg string) {
	if len(msg) != 0 {
		fmt.Printf("ERROR: %v\n", msg)
		os.Exit(1)
	}
	os.Exit(0)
}
