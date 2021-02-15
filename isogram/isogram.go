package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram is gettinga String and return boolean
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	charsCount := map[rune]bool{}
	for _, c := range word {
		if !unicode.IsLetter(c) {
			continue
		}
		if charsCount[c] {
			return false
		}
		charsCount[c] = true
	}
	return true
}
