package robinhood

import (
	"math"
)

func Round(num float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(num*shift+0.5) / shift
}
