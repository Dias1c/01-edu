package main

import (
	"os"
	"tetrisoptimizer/tetrisoptimizer"
)

func main() {
	tetrisoptimizer.Program(os.Args[1:])
}
