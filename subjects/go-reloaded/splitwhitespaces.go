package student

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	word := ""
	for _, symbol := range s {
		if _IsWhiteSpace(symbol) {
			result = append(result, word)
			word = ""
		} else {
			word += string(symbol)
		}
	}
	result = append(result, word)
	word = ""

	return result
}

func _IsWhiteSpace(r rune) bool {
	if r == ' ' || r == '\t' || r == '\n' {
		return true
	}
	return false
}
