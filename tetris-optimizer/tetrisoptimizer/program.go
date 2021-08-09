package tetrisoptimizer

import (
	"fmt"
	"os"
)

func Program(args []string) {
	_CheckForCountArgs(args)
	filepath := args[0]
	fmt.Printf("FilePath: %v\n", filepath)
}

func _CheckForCountArgs(args []string) {
	if len(args) != 1 {
		_CloseProgram("Count arguments != 1;")
	}
}

func _CloseProgram(msg string) {
	if len(msg) != 0 {
		fmt.Printf("ERROR: %v\n", msg)
		os.Exit(1)
	}
	os.Exit(0)
}
