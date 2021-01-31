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
	dr := []rune(d)
	var h = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nucleotide := range dr {
		if nucleotide != 'A' && nucleotide != 'C' && nucleotide != 'G' && nucleotide != 'T' {
			return h, fmt.Errorf("Invaliv nucleotide: %q", nucleotide)
		}
		if nucleotide == 'A' {
			h['A']++
		} else if nucleotide == 'C' {
			h['C']++
		} else if nucleotide == 'G' {
			h['G']++
		} else if nucleotide == 'T' {
			h['T']++
		}
	}
	return h, nil
}
