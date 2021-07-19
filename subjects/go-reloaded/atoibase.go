package student

func AtoiBase(s string, base string) int {
	if !atoiBaseIsBaseCorrect(base) {
		return 0
	}
	numberLen, system := len(s), len(base)
	// Set Digits
	digits := make(map[rune]int, numberLen)
	for idx, d := range base {
		digits[d] = idx
	}

	result := 0
	for _, d := range s {
		result = result*system + digits[d]
	}
	return result
}

func atoiBaseIsBaseCorrect(base string) bool {
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
