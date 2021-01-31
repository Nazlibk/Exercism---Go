package raindrops

import "strconv"

func Convert(num int) string {
	rainDrop := ""
	flag := false
	if num%3 == 0 {
		rainDrop += "Pling"
		flag = true
	}
	if num%5 == 0 {
		rainDrop += "Plang"
		flag = true
	}
	if num%7 == 0 {
		rainDrop += "Plong"
		flag = true
	}
	if !flag {
		rainDrop += strconv.Itoa(num)
	}
	return rainDrop
}
