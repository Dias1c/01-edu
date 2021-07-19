package student

func ConvertBase(nbr, baseFrom, baseTo string) string {
	if !convertBaseIsBaseCorrect(baseFrom) || !convertBaseIsBaseCorrect(baseTo) || convertBaseIsContains(nbr, '-') {
		return ""
	}
	// From
	system := len(baseFrom)
	digits := make(map[rune]int, system)
	for idx, digit := range baseFrom {
		digits[digit] = idx
	}

	numberDec := 0
	for _, digit := range nbr {
		numberDec = numberDec*system + digits[digit]
	}
	// To
	system = len(baseTo)
	result := ""
	for numberDec > 0 {
		digit := numberDec % system
		result = string(baseTo[digit]) + result
		numberDec = numberDec / system
	}
	return result
}

func convertBaseIsBaseCorrect(base string) bool {
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

func convertBaseIsContains(text string, item rune) bool {
	for _, symbol := range text {
		if item == symbol {
			return true
		}
	}
	return false
}
