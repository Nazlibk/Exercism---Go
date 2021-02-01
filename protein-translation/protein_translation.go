package protein

import (
	"errors"
)

var ErrStop error = errors.New("ErrStop")
var ErrInvalidBase error = errors.New("ErrInvalidBase")

func FromCodon(condon string) (string, error) {
	switch condon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

func FromRNA(rna string) ([]string, error) {
	result := make([]string, 0, len(rna)/3)
	var err error
	for i := 0; i < len(rna); i = i + 3 {
		protein, e := FromCodon(rna[i : i+3])
		if e == nil {
			result = append(result, protein)
		} else {
			err = e
			break
		}
	}
	return result, err
}
