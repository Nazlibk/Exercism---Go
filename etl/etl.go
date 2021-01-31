package etl

import "strings"

func Transform(oldMap map[int][]string) map[string]int {
	newMap := make(map[string]int)
	for score, chars := range oldMap {
		for _, char := range chars {
			newMap[strings.ToLower(char)] = score
		}
	}
	return newMap
}
