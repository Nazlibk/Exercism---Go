package anagram

import (
	"strings"
)

func Detect(sub string, list []string) []string {
	result := make([]string, 0)
	sub = strings.ToLower(sub)
	for _, word := range list {
		if len(word) != len(sub) {
			continue
		} else {
			wordLower := strings.ToLower(word)
			if sub == wordLower {
				continue
			}
			for _, char := range sub {
				if strings.ContainsRune(wordLower, char) {
					wordLower = strings.Replace(wordLower, string(char), "", 1)
				}
			}
			if wordLower == "" {
				result = append(result, word)
			}
		}
	}
	return result
}
