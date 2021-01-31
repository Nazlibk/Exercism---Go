package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	// var h Histogram
	dr := []rune(d)
	a := []rune("A")[0]
	c := []rune("C")[0]
	g := []rune("G")[0]
	t := []rune("T")[0]
	var h = Histogram{a: 0, c: 0, g: 0, t: 0}
	for _, nucleotide := range dr {
		if nucleotide != a && nucleotide != c && nucleotide != g && nucleotide != t {
			return h, fmt.Errorf("Invaliv nucleotide: %q", nucleotide)
		}
		if nucleotide == a {
			h[a]++
		} else if nucleotide == c {
			h[c]++
		} else if nucleotide == g {
			h[g]++
		} else if nucleotide == t {
			h[t]++
		}
	}
	return h, nil
}
