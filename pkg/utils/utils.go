package utils

import (
	"math"
)

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// calculateDistance - f
func CalculateEuclidDistance(x, y Vector3) float64 {
	var distance float64
	distance = math.Sqrt(math.Pow((x.X-y.X), 2) + math.Pow((x.Y-y.Y), 2) + math.Pow((x.Z-y.Z), 2))
	return distance
}
