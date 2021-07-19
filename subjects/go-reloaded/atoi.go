package student

// String To Number
func Atoi(s string) int {
	result, size := 0, len(s)
	for idx, element := range s {
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
				runeNum := s[i]
				if '0' <= runeNum && runeNum <= '9' {
					prev = result
					result = result*10 + int(runeNum-48)
					if prev > result {
						break
					}
				} else {
					return 0
				}
				if result < 0 {
					break
				}
			}
			if isNegative {
				result *= -1
				if result > 0 {
					return -9223372036854775808
				}
			} else if result < 0 {
				return 9223372036854775807
			}
			break
		} else {
			return 0
		}
	}
	return result
}
