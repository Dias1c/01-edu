package student

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if !printNbrBaseIsBaseCorrect(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	baseInRunes := []rune(base)
	if nbr == 0 {
		z01.PrintRune(baseInRunes[0])
		return
	}
	system := len(base)
	isNegative := false
	reminder := nbr % system
	nbr = nbr / system
	if nbr < 0 {
		isNegative = true
		nbr, reminder = nbr*-1, reminder*-1
	}
	reversedIndexes := []int{reminder}
	for nbr > 0 {
		reminder = nbr % system
		reversedIndexes = append(reversedIndexes, reminder)
		nbr = nbr / system
	}
	if isNegative {
		z01.PrintRune('-')
	}
	for i := len(reversedIndexes) - 1; i > -1; i-- {
		z01.PrintRune(baseInRunes[reversedIndexes[i]])
	}
}

func printNbrBaseIsBaseCorrect(base string) bool {
	if len(base) < 2 {
		return false
	}
	chars := make(map[rune]int)
	for _, element := range base {
		chars[element]++
		if chars[element] > 1 || element == '-' || element == '+' {
			return false
		}
	}
	return true
}
