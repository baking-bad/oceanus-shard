package system

import (
	"math"
	"oceanus-shard/constants"

	"golang.org/x/exp/rand"
)

func getRandomPointOnCircle(x, y, r float64) [2]float64 {
	angle := rand.Float64() * constants.TwoPi
	var points [2]float64
	points[0] = x + r*math.Cos(angle)
	points[1] = y + r*math.Sin(angle)
	return points
}
