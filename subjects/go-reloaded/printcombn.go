package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n < 1 || n > 9 {
		return
	}
	numbers := make([]rune, n)
	for i := 0; i < n; i++ {
		numbers[i] = 48
	}
	_BackTracking(0, n, numbers)
}

func _BackTracking(idx, size int, numbers []rune) {
	if size == idx {
		_PrintAllRunes(numbers)
		if numbers[0] != rune(58-size) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		} else {
			z01.PrintRune('\n')
		}
		return
	}
	if idx > 0 {
		numbers[idx] = numbers[idx-1] + 1
	}
	for i := numbers[idx]; i <= '9'; i++ {
		_BackTracking(idx+1, size, numbers)
		numbers[idx]++
		if numbers[idx] > '9' {
			break
		}
	}
}

func _PrintAllRunes(runes []rune) {
	for _, element := range runes {
		z01.PrintRune(element)
	}
}
