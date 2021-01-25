package acronym

import "strings"

func Abbreviate(s string) string {
	s1 := strings.ReplaceAll(s, "-", " ")
	s2 := strings.ReplaceAll(s1, "_", " ")
	words := strings.Split(s2, " ")
	abb := ""
	for _, w := range words {
		if w != "" {
			abb += strings.ToUpper(string(w[0]))
		}
	}
	return abb
}
