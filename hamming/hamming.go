package hamming

import "fmt"

// Distance is a function to calculate distance between two strings
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, fmt.Errorf("unequal lengths %q, %q", a, b)
	}
	count := 0
	for i := range ar {
		if ar[i] != br[i] {
			count++
		}
	}
	return count, nil
}
