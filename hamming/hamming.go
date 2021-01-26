package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) != len(b) {
		return 0, errors.New("length of two string is different")
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count = count + 1
			}
		}
	}
	return count, nil
}
