package collatzconjecture

import "errors"

func CollatzConjecture(input int) (int, error) {
	
	steps := 0

	if input <= 0 {
		return 0, errors.New("input cannot be zero")
	}

	for input > 1 {
		if input % 2 == 0 {
			input = input / 2
		} else {
			input = (input * 3) + 1
		}
		
		steps++
	}
	
	return steps, nil
}
