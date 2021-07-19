package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fileNames := os.Args[1:]
	if len(fileNames) == 0 {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			panic(err)
		}
		return
	}
	for _, fName := range fileNames {
		bytes, err := ioutil.ReadFile(fName)
		if err != nil {
			PrintString("ERROR: " + fName + ": No such file or directory\n")
		} else {
			PrintString(string(bytes))
		}
	}
}

func PrintString(data string) {
	for _, s := range data {
		z01.PrintRune(s)
	}
}
