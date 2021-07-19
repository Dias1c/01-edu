package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	size := len(args)
	if size == 0 {
		os.Exit(1)
	}
	isHaveParam, cnt, errMessage := false, uint64(0), ""
	args, isHaveParam, cnt, errMessage = RemoveAndGetParametr(args)
	if errMessage != "" {
		fmt.Printf("%v", errMessage)
		os.Exit(1)
	}
	if isHaveParam {
		size := len(args)
		if size == 0 {
			os.Exit(1)
		}
		isHadProblem := false
		for idx, fName := range args {
			bytes, err := os.ReadFile(fName)
			if err != nil {
				fmt.Printf("tail: cannot open '%v' for reading: No such file or directory\n", fName)
				isHadProblem = true
				continue
			}
			bytesLen := uint64(len(bytes))
			start := bytesLen - cnt
			if start > 0 {
				start = 0
			}
			if size > 1 {
				if idx > 0 && !isHadProblem {
					fmt.Printf("\n")
				}
				fmt.Printf("==> %v <==\n", fName)
			}
			fmt.Printf("%s", bytes[start:])
			isHadProblem = false
		}
	} else { //Work Like Cat
		isHadProblem := false
		for idx, fName := range args {
			bytes, err := os.ReadFile(fName)
			if err != nil {
				fmt.Printf("tail: cannot open '%v' for reading: No such file or directory\n", fName)
				isHadProblem = true
				continue
			}
			if size > 1 {
				if idx > 0 && !isHadProblem {
					fmt.Printf("\n")
				}
				fmt.Printf("==> %v <==\n", fName)
			}
			fmt.Printf("%s", bytes)
			isHadProblem = false
		}
	}
	os.Exit(0)
}

// args, isHave -c, Num, IsNumCorrect
func RemoveAndGetParametr(args []string) ([]string, bool, uint64, string) {
	size := len(args)
	isHaveParam, num, errorMessage := false, uint64(0), ""
	for i := 0; i < size; i++ {
		if args[i][0] == '-' {
			if len(args[i]) == 1 {
				os.Exit(1)
			} else if args[i] != "-c" {
				return args, true, 0, "tail: invalid option -- '" + string(args[i][1]) + "'\n"
			} else if i == size-1 {
				return args, true, 0, "tail: option requires an argument -- '" + args[i][1:] + "'\n"
			} else {
				nextIdx := i + 1
				if len(args[nextIdx]) > 0 && args[nextIdx][0] == '-' {
					args[nextIdx] = args[nextIdx][1:]
				}
				tNum, isCorrect := StringToUInt64(args[nextIdx])
				if !isCorrect {
					return args, true, 0, "tail: invalid number of bytes: ‘" + args[nextIdx] + "’\n"
				}
				num = tNum
				isHaveParam = true
			}
			args = append(args[:i], args[i+2:]...)
			i--
			size = len(args)
		}
	}
	return args, isHaveParam, num, errorMessage
}

func StringToUInt64(number string) (uint64, bool) {
	result, size := uint64(0), len(number)
	if size == 0 {
		return 0, false
	}
	for idx, element := range number {
		if (idx == 0 && element == '+') || '0' <= element && element <= '9' {
			if element == '+' {
				idx++
			}
			prev := uint64(0)
			for i := idx; i < size; i++ {
				runeNum := number[i]
				if '0' <= runeNum && runeNum <= '9' {
					prev = result
					result = result*10 + uint64(runeNum-48)
					if prev > result {
						return 0, false
					}
				} else {
					return 0, false
				}
			}
			break
		} else {
			return 0, false
		}
	}
	return result, true
}
