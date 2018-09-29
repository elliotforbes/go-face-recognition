package utils

import "math"

type Histogram struct {
	Bin    Bin
	Values []float64
}

type Bin struct {
	High float64
	Low  float64
}

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func GenerateHistogram(values []float64) Histogram {
	const buckets = 10
	var histogram Histogram

	for i := range values {
		histogram.Values = append(histogram.Values, float64(i))
	}

	return histogram
}

// calculateDistance - f
func CalculateEuclidDistance(x, y Vector3) float64 {
	var distance float64

	distance = math.Sqrt(math.Pow((x.X-y.X), 2) + math.Pow((x.Y-y.Y), 2) + math.Pow((x.Z-y.Z), 2))

	return distance
}
