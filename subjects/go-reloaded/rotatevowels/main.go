package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	cntArgs := len(os.Args)
	if cntArgs < 2 {
		z01.PrintRune('\n')
		return
	}
	word := []rune(os.Args[1])
	for i := 2; i < cntArgs; i++ {
		word = append(word, ' ')
		word = append(word, []rune(os.Args[i])...)
	}
	for l, r := 0, len(word)-1; l < r; {
		lIsVowel, rIsVowel := isVowel(word[l]), isVowel(word[r])
		if !lIsVowel {
			l++
		}
		if !rIsVowel {
			r--
		}
		if lIsVowel && rIsVowel {
			word[l], word[r] = word[r], word[l]
			l, r = l+1, r-1
		}
	}
	PrintString(string(word) + "\n")
}

func isVowel(symbol rune) bool {
	if 'A' <= symbol && symbol <= 'Z' {
		symbol = symbol + 32
	}
	if symbol == 'a' || symbol == 'e' || symbol == 'i' || symbol == 'o' || symbol == 'u' {
		return true
	}
	return false
}

func PrintString(word string) {
	for _, element := range word {
		z01.PrintRune(element)
	}
}
