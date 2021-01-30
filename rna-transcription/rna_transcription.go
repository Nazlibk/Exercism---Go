package strand

func ToRNA(dna string) string {
	dnar := []rune(dna)
	for i, d := range dnar {
		if d == []rune("G")[0] {
			dnar[i] = []rune("C")[0]
		} else if d == []rune("C")[0] {
			dnar[i] = []rune("G")[0]
		} else if d == []rune("T")[0] {
			dnar[i] = []rune("A")[0]
		} else if d == []rune("A")[0] {
			dnar[i] = []rune("U")[0]
		}
	}
	return string(dnar)
}
