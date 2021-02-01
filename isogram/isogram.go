package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	charsCount := make(map[rune]int)
	wordr := []rune(word)
	for _, c := range wordr {
		if unicode.IsLetter(c) {
			if _, ok := charsCount[c]; ok {
				return false
			}
			charsCount[c]++

		}
	}
	return true
}
