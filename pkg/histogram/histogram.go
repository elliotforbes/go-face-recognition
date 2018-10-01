package histogram

type Histogram struct {
	Values []float64
}

// GenerateHistogram generates a histogram based off the pixel 2D array passed into it
func GenerateHistogram(pixels [][]float64) Histogram {
	var histogram Histogram

	// 1 generate grid

	return histogram
}

// EuclidDistance returns the euclidean distance between two histograms
func (h *Histogram) EuclidDistance(histogram Histogram) (float64, error) {
	var distance float64
	// distance = math.Sqrt(math.Pow((x.X-y.X), 2) + math.Pow((x.Y-y.Y), 2) + math.Pow((x.Z-y.Z), 2))
	// return 0, distance
	return distance, nil
}

// NormalizedEuclidDistance returns the normalized euclidean distance between two histograms
func (h *Histogram) NormalizedEuclidDistance(histogram Histogram) (float64, error) {
	var distance float64
	return distance, nil
}

// AbsoluteValue returns the absolute value between two histograms
func (h *Histogram) AbsoluteValue(histogram Histogram) (float64, error) {
	var value float64
	return value, nil
}

// ChiSquare returns the distance between two histograms calculated using the chi square statistic
func (h *Histogram) ChiSquare(histogram Histogram) (float64, error) {
	var sum float64
	return sum, nil
}
