package luhn

import (
	"strings"
	"unicode"
)

// Valid is a method that validate if num is ISN
func Valid(num string) bool {
	num = strings.ReplaceAll(num, " ", "")
	if len(num) <= 1 {
		return false
	}
	if len(num)%2 != 0 {
		num = "0" + num
	}
	sum := 0
	for i, n := range num {
		intN := int(n - '0')
		if !unicode.IsDigit(n) {
			return false
		}
		if (i+1)%2 == 0 {
			sum += intN
		} else {
			multi := intN * 2
			if multi > 9 {
				sum += multi - 9
			} else {
				sum += multi
			}
		}
	}
	if sum%10 == 0 {
		return true
	}
	return false
}
