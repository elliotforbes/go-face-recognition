package math

import (
	"math"

	"github.com/elliotforbes/go-face-recognition/pkg/histogram"
)

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// calculateDistance - f
func (h *histogram.Histogram) EuclidDistance(histogram histogram.Histogram) (float64, error) {
	var distance float64
	distance = math.Sqrt(math.Pow((x.X-y.X), 2) + math.Pow((x.Y-y.Y), 2) + math.Pow((x.Z-y.Z), 2))
	return 0, distance
}
