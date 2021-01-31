package accumulate

func Accumulate(words []string, accFunc func(string) string) []string {
	result := make([]string, 0, len(words))
	for _, w := range words {
		result = append(result, accFunc(w))
	}
	return result
}
