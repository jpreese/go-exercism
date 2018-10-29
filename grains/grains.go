package grains

import (
	"errors"
)

// Square returns the number of grains on that square
func Square(number int) (uint64, error) {

	if number <= 0 || number > 64 {
		return 0, errors.New("invalid argument")
	}

	result := uint64(1 << uint(number-1))

	return result, nil
}

// Total returns the total number of grains
func Total() uint64 {
	return (1 << 64) - 1
}
