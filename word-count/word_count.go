package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(sentence string) Frequency {
	sentence = strings.ToLower(sentence)
	sentence = strings.ReplaceAll(sentence, "\n", "")
	sentence = strings.ReplaceAll(sentence, "\t", "")
	sentence = strings.ReplaceAll(sentence, ",", " ")
	sentence = strings.Trim(sentence, "!&@$%^ ")
	singleSpacePattern := regexp.MustCompile(`\s+`)
	sentence = singleSpacePattern.ReplaceAllString(sentence, " ")

	words := strings.Split(sentence, " ")
	freq := map[string]int{}
	for _, w := range words {
		w = strings.Trim(w, "'.:")
		freq[w]++
	}
	return freq
}
