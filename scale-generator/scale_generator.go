package scale

func Scale(tonic string, interval string) []string {
	sharp := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flat := []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	noSharpNoFlast := []string{"C", "a"}
	useSharp := []string{"G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#"}
	useFlat := []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}
	var result []string
	if contains(noSharpNoFlast, tonic) {
		result = getMusicalScale(sharp, interval, tonic)
	} else if contains(useSharp, tonic) {
		result = getMusicalScale(sharp, interval, tonic)
	} else if contains(useFlat, tonic) {
		result = getMusicalScale(flat, interval, tonic)
	}
	return result
}

func getMusicalScale(toneList []string, interval string, tonic string) []string {
	var result []string
	index := 0
	for i := 0; i < len(toneList); i++ {
		if toneList[i] == tonic {
			index = i
		}
	}
	if interval == "" {
		for count := 0; count <= 12; count++ {
			result[count] = toneList[index]
			index++
			if index >= len(toneList) {
				index = 0
			}
		}
	} else {
		for count := 0; count <= 7; count++ {
			step := 1
			if string(interval[index]) == "M" {
				step = 2
			}
			result[count] = toneList[index]
			index = index + step
			if index >= len(toneList) {
				index = index - len(toneList)
			}
		}
	}
	return result
}

func contains(list []string, tonic string) bool {
	for _, a := range list {
		if a == tonic {
			return true
		}
	}
	return false
}
