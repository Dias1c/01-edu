package asciiart

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var GConfigs = Configs{Theme: "standard"}

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
	// Import Arguments into GConfig
	received, imported, err, hint := ImportParams(params)
	if err != nil {
		CloseProgramAndHint(err, hint)
	} else if GConfigs.ReverseArg && imported {
		CloseProgram(ErrSingleReverse)
	} else if !imported && !GConfigs.ReverseArg {
		CloseProgram(ErrNoReceiving)
	} else if GConfigs.Conflict() {
		CloseProgramAndHint(ErrConflictArgs, GConfigs.ConflictHint())
	}
	// Import Banner
	err = ImportBanner()
	if err != nil {
		CloseProgram(err)
	}
	// todo Reverse
	if GConfigs.ReverseArg {
		CloseProgram(nil)
	}
	// Get Graphic Representation of Received String
	received = GetFormattedString(received)
	gr, err, hint := GetGraphicRepresentation(received)
	if err != nil {
		CloseProgramAndHint(err, hint)
	}
	// Output to File
	if GConfigs.OutputArg {
		file, err := os.OpenFile(GConfigs.Output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
		if err != nil {
			CloseProgram(err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		for i := range gr {
			for j := range gr[i] {
				_, err = writer.WriteString(gr[i][j])
				if err != nil {
					CloseProgram(err)
				}
				writer.WriteRune('\n')
			}
		}
		err = writer.Flush()
		CloseProgram(err)
	}
	// Print in Terminal
	// to do with color and align
	PrintGraphicRepresentation(gr)
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

// Close Program with Hint
func CloseProgramAndHint(err error, hint string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println(hint)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

// Imports Params Into GConfigs and Returns Received Strings
func ImportParams(params []string) (string, bool, error, string) {
	var received string
	var imported bool
	paramRegexp := regexp.MustCompile(`--(?s:.*)`)
	optionRegexp := regexp.MustCompile(`--(?P<name>[[:lower:]]{5,7})=(?P<option>(?s:.*))`)
	for i := 0; i < len(params); i++ {
		if paramRegexp.MatchString(params[i]) {
			if optionRegexp.MatchString(params[i]) {
				matches := optionRegexp.FindStringSubmatch(params[i])
				nameInd, optionInd := optionRegexp.SubexpIndex("name"), optionRegexp.SubexpIndex("option")
				name, option := matches[nameInd], matches[optionInd]
				switch name {
				case "color":
					err := GConfigs.ImportColor(option)
					if err != nil {
						return "", false, err, option
					}
					break
				case "output":
					err := GConfigs.ImportOutput(option)
					if err != nil {
						return "", false, err, option
					}
					break
				case "reverse":
					err := GConfigs.ImportReverse(option)
					if err != nil {
						return "", false, err, option
					}
					break
				default:
					return "", false, ErrNotHandleArg, name
				}
			} else {
				return "", false, ErrArgument, params[i]
			}
		} else {
			if GConfigs.ThemeArg {
				if imported {
					return "", false, ErrManyReceiving, params[i]
				} else {
					received = params[i]
					imported = true
				}
			} else {
				if !GConfigs.ImportTheme(params[i]) {
					if imported {
						return "", false, ErrManyReceiving, params[i]
					} else {
						received = params[i]
						imported = true
					}
				}
			}
		}
	}
	return received, imported, nil, ""
}
