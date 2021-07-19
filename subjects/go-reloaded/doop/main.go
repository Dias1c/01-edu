package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	numStr1 := os.Args[1]
	oper := os.Args[2]
	numStr2 := os.Args[3]
	a, err1 := StringToInt(numStr1)
	b, err2 := StringToInt(numStr2)
	if err1 != "" || err2 != "" {
		PrintString("0\n")
		return
	}
	res := ""
	switch oper {
	case "+":
		res = Plus(a, b)
	case "-":
		res = Minus(a, b)
	case "*":
		res = Mult(a, b)
	case "/":
		res = Del(a, b)
	case "%":
		res = Mod(a, b)
	default:
		res = "0"
	}
	PrintString(res + "\n")
}

func Mult(a, b int) string {
	res := a * b
	aLen := GetNumLength(a)
	blen := GetNumLength(b)
	reslen := GetNumLength(res)
	if (a < 0 && b > 0 || a > 0 && b < 0) && (res > 0 || res > a || res > b) {
		return "0"
	} else if a > 0 && b > 0 && (res <= 0 || res < a || res < b) {
		return "0"
	} else if a < 0 && b < 0 && res < 0 {
		return "0"
	} else if ((aLen > 1 && blen >= 1 && b != 1 && b != -1) || (aLen >= 1 && blen > 1 && a != 1 && a != -1)) && (reslen <= aLen || reslen <= blen) {
		return "0"
	}
	return IntToString(res)
}

func Del(a, b int) string {
	if b == 0 {
		return "No division by 0"
	}
	res := a / b
	if b < 0 && a < 0 && res < 0 {
		return "0"
	}
	return IntToString(res)
}

func Plus(a, b int) string {
	res := a + b
	if a > 0 && b > 0 && (res <= 0 || res < a || res < b) {
		return "0"
	} else if a < 0 && b < 0 && (res > 0 || res > a || res > b) {
		return "0"
	}
	return IntToString(res)
}

func Minus(a, b int) string {
	res := a - b
	if a == b || a < 0 && b < 0 && (res > a || res > b) {
		return "0"
	} else if a > 0 && b < 0 && res < a {
		return "0"
	} else if a < 0 && b > 0 && res > a {
		return "0"
	}
	return IntToString(res)
}

func Mod(a, b int) string {
	if b == 0 {
		return "No modulo by 0"
	}
	res := a % b
	return IntToString(res)
}

func PrintString(str string) {
	for _, element := range str {
		z01.PrintRune(element)
	}
}

func StringToInt(number string) (int, string) {
	result, size := 0, len(number)
	for idx, element := range number {
		if element == '-' || element == '+' || '0' <= element && element <= '9' {
			isNegative := false
			if element == '-' {
				isNegative = true
				idx++
			} else if element == '+' {
				idx++
			}
			prev := result
			for i := idx; i < size; i++ {
				runeNum := number[i]
				if '0' <= runeNum && runeNum <= '9' {
					prev = result
					result = result*10 + int(runeNum-48)
					if prev > result {
						return 0, "err"
					}
				} else {
					return 0, "err"
				}
			}
			if isNegative {
				result *= -1
				if result > 0 {
					return 0, "err"
				}
			} else if result < 0 {
				return 0, "err"
			}
			break
		} else {
			return 0, "err"
		}
	}
	return result, ""
}

func IntToString(num int) string {
	res := ""
	isNegative := false
	if num < 0 {
		isNegative = true
	} else if num == 0 {
		return "0"
	}
	for num != 0 {
		digit := num % 10
		if digit < 0 {
			digit *= -1
		}
		res = string(digit+48) + res
		num = num / 10
	}
	if isNegative {
		return "-" + res
	}
	return res
}

// Gets NumLength, FirstNum
func GetNumLength(num int) int {
	if num == 0 {
		return 1
	}
	cnt := 0
	for num != 0 {
		cnt++
		num = num / 10
	}
	return cnt
}
