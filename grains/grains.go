package grains

import (
	"errors"
	"math/big"
)

// Square returns the number of grains on that square
func Square(number int) (uint64, error) {

	if number <= 0 || number > 64 {
		return 0, errors.New("invalid argument")
	}

	result := (new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(number-1)), nil).Uint64())
	return result, nil
}

// Total returns the total number of grains
func Total() uint64 {
	currentSquare := uint64(0)
	sum := uint64(0)

	for i := 1; i <= 64; i++ {
		currentSquare, _ = Square(i)
		sum += currentSquare
	}

	return sum
}
