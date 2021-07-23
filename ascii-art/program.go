package asciiart

import (
	"fmt"
	"os"
	"regexp"
	"strings"
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
	received, imported, err, hint := ImportParams(params)
	if err != nil {
		CloseProgramAndHint(err, hint)
	} else if GConfigs.ReverseArg && imported {
		CloseProgram(ErrSingleReverse)
	} else if !imported && !GConfigs.ReverseArg {
		CloseProgram(ErrNoReceiving)
	}
	err = ImportBanner()
	if err != nil {
		CloseProgram(err)
	}
	received = GetFormattedString(received)
	PrintWithBanner(received)
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

func CloseProgramAndHint(err error, hint string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println(hint)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

// Add New Line in String
func GetFormattedString(str string) string {
	return strings.ReplaceAll(str, "\\n", "\n")
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
