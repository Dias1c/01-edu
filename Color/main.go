package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	CSI       = "\033["
	F256      = "38"
	Bg256     = "48"
	FgDefault = "39"
	BgDefault = "49"
)

func newColorRGB(str string, red, green, blue, red1, green1, blue1 uint8, coloris bool) string {
	if coloris == true {
		ired := strconv.Itoa(int(red))
		igreen := strconv.Itoa(int(green))
		iblue := strconv.Itoa(int(blue))
		tstr := CSI + F256 + ";2;" + ired + ";" + igreen + ";" + iblue + "m"
		tstr += str
		return tstr + CSI + FgDefault + ";5;" + BgDefault + ";5;" + "m"
	}
	ired1 := strconv.Itoa(int(red1))
	igreen1 := strconv.Itoa(int(green1))
	iblue1 := strconv.Itoa(int(blue1))
	tstr1 := CSI + F256 + ";2;" + ired1 + ";" + igreen1 + ";" + iblue1 + "m"
	tstr1 += str
	return tstr1 + CSI + FgDefault + ";5;" + BgDefault + ";5;" + "m"
}

func main() {
	args := os.Args[1:]
	if cnt := len(args); cnt != 4 {
		fmt.Printf("Not Enough args %d != 4;\nTry: go run main.go word red green blue\n", cnt)
		os.Exit(1)
	}
	word, red, green, blue := args[0], args[1], args[2], args[3]
	fmt.Printf("\033[%s%s%s]m%s\n", red, green, blue, word)
	// word := args[0]
	// for r := 0; r < 256; r++ {
	// 	for g := 0; g < 256; g++ {
	// 		for b := 0; b < 256; b++ {
	// 			red := getNumStr3(r)
	// 			green := getNumStr3(g)
	// 			blue := getNumStr3(b)
	// 			fmt.Printf("\033[%v%v%vm%s\033[0m R: %d, G: %d, B: %d\n", red, green, blue, word, r, g, b)
	// 		}
	// 	}
	// }
}

func getNumStr3(num int) string {
	if num < 10 {
		return fmt.Sprintf("00%v", num)
	} else if num < 100 {
		return fmt.Sprintf("0%v", num)
	} else if num < 1000 {
		return fmt.Sprintf("%v", num)
	}
	return "0"
}
