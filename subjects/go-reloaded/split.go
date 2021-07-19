package student

func Split(s, sep string) []string {
	sizeMain, sizeSep := len(s), len(sep)
	if sizeMain < sizeSep {
		return []string{}
	} else if sizeSep == 0 {
		symbols := []string{}
		for _, element := range s {
			symbols = append(symbols, string(element))
		}
		return symbols
	}
	mainWord := []rune(s)
	result, word := []string{}, ""
	for i := 0; i < sizeMain; i++ {
		if i <= sizeMain-sizeSep && splitIsContains(string(mainWord[i:i+sizeSep]), sep) {
			result = append(result, word)
			i = i + sizeSep - 1
			word = ""
		} else {
			word += string(mainWord[i])
		}
	}
	result = append(result, word)
	word = ""
	return result
}

func splitIsContains(orig, consist string) bool {
	sizeOrig, sizeConsist := len(orig), len(consist)
	if sizeOrig != sizeConsist {
		return false
	}
	for i := 0; i < sizeOrig; i++ {
		if orig[i] != consist[i] {
			return false
		}
	}
	return true
}
