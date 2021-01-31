package romannumerals

import "fmt"

func ToRomanNumeral(num int) (string, error) {
	if num <= 0 || num > 3000 {
		return "", fmt.Errorf("Numebr %q is not valid", num)
	}
	romeNum := ""
	maps := map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M"}
	devider := 1000
	for devider != 0 {
		n := num / devider
		num = num % devider
		if n < 4 {
			for i := 1; i <= n; i++ {
				romeNum += maps[devider]
			}
		} else if n == 4 {
			romeNum += maps[devider]
			romeNum += maps[5*devider]
		} else if n == 5 {
			romeNum += maps[5*devider]
		} else if n > 5 && n < 9 {
			romeNum += maps[5*devider]
			for i := 6; i <= n; i++ {
				romeNum += maps[devider]
			}
		} else if n == 9 {
			romeNum += maps[devider/10]
			romeNum += maps[devider]
		}
		devider = devider / 10
	}
	return romeNum, nil
}
