package asciiart

import (
	"fmt"
	"os"
	"strings"
)

var GConfigs = &Configs{FileTheme: "standard.txt", Color: "white"}
var GThemes []string

// var GColors = map[string]string{}

// func InitColors() {
// 	GColors := make(map[string]string, 10)
// 	GColors["reset"] = "\033[0m"
// 	GColors["red"] = "\033[31m"
// 	GColors["green"] = "\033[32m"
// 	GColors["yellow"] = "\033[33m"
// 	GColors["blue"] = "\033[34m"
// 	GColors["purple"] = "\033[35m"
// 	GColors["cyan"] = "\033[36m"
// 	GColors["gray"] = "\033[37m"
// 	GColors["white"] = "\033[97m"
// }

func Program(params []string) {
	length := len(params)
	if length == 0 {
		CloseProgram(nil)
	}
	// InitColors()
	ReadConfigs()
	ImportBanner()
	output := GetFormattedString(params[0])
	PrintWithBanner(output)
}

// Close Program with Error if err != nil
func CloseProgram(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func GetFormattedString(str string) string {
	return strings.ReplaceAll(str, "\\n", "\n")
}
