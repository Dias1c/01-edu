package main

import (
	"fmt"
	"os"
	"regexp"
)

const (
	CSI      = "\033["
	FgRGB256 = "38"
	// Bg256     = "48"
	// FgDefault = "39"
	// BgDefault = "49"
)

// func newColorRGB(str string, red, green, blue, red1, green1, blue1 uint8, coloris bool) string {
// 	if coloris == true { // FG
// 		ired := strconv.Itoa(int(red))
// 		igreen := strconv.Itoa(int(green))
// 		iblue := strconv.Itoa(int(blue))
// 		tstr := CSI + Fg256 + ";2;" + ired + ";" + igreen + ";" + iblue + "m"
// 		tstr += str
// 		return tstr + CSI + FgDefault + ";5;" + BgDefault + ";5;" + "m"
// 	}
// 	ired1 := strconv.Itoa(int(red1))
// 	igreen1 := strconv.Itoa(int(green1))
// 	iblue1 := strconv.Itoa(int(blue1))
// 	tstr1 := CSI + Fg256 + ";2;" + ired1 + ";" + igreen1 + ";" + iblue1 + "m"
// 	tstr1 += str
// 	return tstr1 + CSI + FgDefault + ";5;" + BgDefault + ";5;" + "m"
// }

func main() {
	args := os.Args[1:]
	if cnt := len(args); cnt != 1 {
		fmt.Printf("Not Enough args %d != 1;\n%q\n", cnt, args)
		os.Exit(1)
	}
	colors = make(map[string]string, 0)
	code := getColorCode(args[0])
	fmt.Println(code, "Hello World")
}

func getColor(red, green, blue string) string {
	return fmt.Sprintf("\033[38;2;%s;%s;%sm", red, green, blue)
}

var colors map[string]string

func getColorCode(colour string) string {
	// CSI       = "\033["
	// FgRGB256  = "38;2;"
	// BgRGB256  = "48"
	// FgDefault = "39"
	// BgDefault = "49"
	if code, isIn := colors[colour]; isIn {
		return code
	}
	pattern := regexp.MustCompile(`^RGB\(([\d]{1,3}),([\d]{1,3}),([\d]{1,3})\)$`)
	isMatched := pattern.MatchString(colour)
	if !isMatched {
		err := fmt.Errorf("Uncorrect color value: '%s'!", colour)
		_CloseProgram(err)
	}
	subMatchs := pattern.FindStringSubmatch(colour)
	red, green, blue := subMatchs[1], subMatchs[2], subMatchs[3]

	code := fmt.Sprintf("\033[38;2;%s;%s;%sm", red, green, blue)
	colors[colour] = code
	return code
}

func _CloseProgram(err error) {
	if err != nil {
		fmt.Printf("%sERROR: %s%s\n", "\033[31m", err.Error(), "\033[0m")
		os.Exit(1)
	}
	os.Exit(0)
}
