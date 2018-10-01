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
func GenerateHistogram(buckets []Bucket, values []float64) Histogram {
	var histogram Histogram
	histogram.Buckets = buckets

	for _, i := range values {
		for j := 0; j < len(histogram.Buckets); j++ {
			if (i > histogram.Buckets[j].Low) && (i < histogram.Buckets[j].High) {
				histogram.Buckets[j].Values = append(histogram.Buckets[j].Values, i)
				break
			}
		}
	}

	return histogram
}
