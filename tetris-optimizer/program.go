package tetrisoptimizer

import (
	"errors"
	"fmt"
	"os"
)

func Program(args []string) {
	_CheckForCountArgs(args)
	filePath := args[0]
	tetrimonies := _GetTetrymoniesFromFile(filePath)
	result := _MatchMinSquare(tetrimonies)
	_ShowResult(result)
}

func _CheckForCountArgs(args []string) {
	if cnt := len(args); cnt != 1 {
		msg := fmt.Sprintf("Program get only 1 param (fileName). But you input %v params", cnt)
		CloseProgram(errors.New(msg))
	}
}

func CloseProgram(err error) {
	if err != nil {
		fmt.Printf("\033[1;31mERROR: %s\033[0m\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
