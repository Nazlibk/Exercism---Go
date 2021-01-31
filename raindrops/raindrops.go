package raindrops

import "strconv"

func Convert(num int) string {
	rainDrop := ""
	if num%3 == 0 {
		rainDrop += "Pling"
	}
	if num%5 == 0 {
		rainDrop += "Plang"
	}
	if num%7 == 0 {
		rainDrop += "Plong"
	}
	if rainDrop == "" {
		rainDrop += strconv.Itoa(num)
	}
	return rainDrop
}
