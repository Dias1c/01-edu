package main

import (
	"asciiart"
	"os"
)

func main() {
	length := len(os.Args)
	if length < 2 {
		asciiart.CloseProgram(asciiart.ErrNoArgument)
	}
	asciiart.Program(os.Args[1:])
}
