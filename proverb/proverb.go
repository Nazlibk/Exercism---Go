package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	proverbs := make([]string, 0, len(rhyme))
	if len(rhyme) == 0 {
		return []string{}
	}
	if len(rhyme) == 1 {
		r := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		proverbs = append(proverbs, r)
		return proverbs
	}
	for i := 0; i < len(rhyme)-1; i++ {
		r := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		proverbs = append(proverbs, r)
	}
	r := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	return append(proverbs, r)
}
