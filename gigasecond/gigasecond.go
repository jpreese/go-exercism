// Contains functionality to add a gigasecond to a given time
package gigasecond

import "time"

// Adds a gigasecond to the provided time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
