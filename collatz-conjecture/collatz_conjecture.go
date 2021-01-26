package collatzconjecture

import "errors"

func CollatzConjecture(input int) (int, error) {
	if input == 0 || input < 0 {
		return input, errors.New("input is not valid")
	}
	steps := 0
	for input != 1 {
		if input%2 == 0 {
			input = input / 2
		} else {
			input = input*3 + 1
		}
		steps = steps + 1
	}
	return steps, nil
}
