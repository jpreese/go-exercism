package suchrandom

import (
	"math/rand"
)

var theMap = map[string]bool{}
var theSlice = []string{}

const iterations = 1000

// RunMap runs maps
func RunMap() bool {

	initDataStructures()

	stringToFind := theSlice[rand.Intn(iterations)]
	_, exists := theMap[stringToFind]
	return exists
}

// RunSlice runes slices
func RunSlice() bool {

	initDataStructures()

	stringToFind := theSlice[rand.Intn(iterations)]

	for _, currentString := range theSlice {
		if stringToFind == currentString {
			return true
		}
	}

	return false
}

func initDataStructures() {

	for i := 0; i < iterations; i++ {
		currentRandomString := randomString(4)

		theMap[currentRandomString] = true
		theSlice = append(theSlice, currentRandomString)
	}

}

func randomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
