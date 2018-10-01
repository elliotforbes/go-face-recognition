package histogram

type Histogram struct {
	Buckets []Bucket
}

type Bucket struct {
	Low    float64
	High   float64
	Values []float64
}

// GenerateHistogram generates a simple histogram in which
// it buckets everything
func GenerateHistogram(pixels [][]float64) Histogram {
	var histogram Histogram

	return histogram
}

// EuclidDistance returns the euclidean distance between two histograms
// or an error
func (h *Histogram) EuclidDistance(histogram Histogram) (float64, error) {
	var distance float64
	// distance = math.Sqrt(math.Pow((x.X-y.X), 2) + math.Pow((x.Y-y.Y), 2) + math.Pow((x.Z-y.Z), 2))
	// return 0, distance
	return distance, nil
}
