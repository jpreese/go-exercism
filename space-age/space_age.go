// Calculates the age of a person (in earth years) for a given planet
package space

const secondsInEarthYear = 31557600.0

type Planet string

var orbitalPeriods = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Given the age of a person (in seconds) on earth and a planet, return
// the age of the person on the given planet
func Age(seconds float64, planet Planet) float64 {

	secondsInYearForPlanet := secondsInEarthYear * orbitalPeriods[planet]
	earthYears := seconds / secondsInYearForPlanet

	return earthYears
}
