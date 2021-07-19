package main

import (
	"asciiart"
	"os"
)

func main() {
	asciiart.Program(os.Args[1:])
}
