package twofer

import "fmt"

func ShareWith(name string) string {
	s := name
	if name == "" {
		s = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", s)
}
