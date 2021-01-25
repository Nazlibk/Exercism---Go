package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = replace(remark)
	if strings.HasSuffix(remark, "?") && strings.Compare(strings.ToUpper(remark), remark) == 0 && hasLetter(remark) {
		return "Calm down, I know what I'm doing!"
	} else if remark == "" {
		return "Fine. Be that way!"
	} else if strings.HasSuffix(remark, "?") {
		return "Sure."
	} else if !hasLetter(remark) {
		return "Whatever."
	} else if strings.Compare(strings.ToUpper(remark), remark) == 0 {
		return "Whoa, chill out!"
	}
	return "Whatever."
}

func hasLetter(s string) bool {
	notLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			notLetter = true
		}
	}
	return notLetter
}

func replace(s string) string {
	s1 := strings.ReplaceAll(s, " ", "")
	s2 := strings.ReplaceAll(s1, "\t", "")
	s3 := strings.ReplaceAll(s2, "\r", "")
	return strings.ReplaceAll(s3, "\n", "")
}
