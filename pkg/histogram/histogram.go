package histogram

import (
	"errors"
	"fmt"
	"math"
)

type Histogram struct {
	Values []uint64
}

// GenerateHistogram generates a histogram based off the pixel 2D array passed into it
func GenerateHistogram(pixels [][]uint64, gridX, gridY uint8) (Histogram, error) {
	var histogram Histogram
	fmt.Println("Generating Histogram")

	if len(pixels) == 0 {
		return histogram, errors.New("Empty Pixels Array")
	}

	rows := len(pixels)
	cols := len(pixels[0])

	if gridX <= 0 || int(gridX) >= cols {
		return histogram, errors.New("Invalid Grid X passed")
	}
	if gridY <= 0 || int(gridY) >= rows {
		return histogram, errors.New("Invalid Grid X passed")
	}

	gridWidth := cols / int(gridX)
	gridHeight := rows / int(gridY)

	// Calculates the histogram of each grid
	for gX := 0; gX < int(gridX); gX++ {
		for gY := 0; gY < int(gridY); gY++ {
			// Create a slice with empty 256 positions
			regionHistogram := make([]uint64, 256)

			// Define the start and end positions for the following loop
			startPosX := gX * gridWidth
			startPosY := gY * gridHeight
			endPosX := (gX + 1) * gridWidth
			endPosY := (gY + 1) * gridHeight

			// Make sure that no pixel has been leave at the end
			if gX == int(gridX)-1 {
				endPosX = cols
			}
			if gY == int(gridY)-1 {
				endPosY = rows
			}

			// Creates the histogram for the current region
			for x := startPosX; x < endPosX; x++ {
				for y := startPosY; y < endPosY; y++ {
					// Make sure we are trying to access a valid position
					if x < len(pixels) {
						if y < len(pixels[x]) {
							if int(pixels[x][y]) < len(regionHistogram) {

								regionHistogram[int(pixels[x][y])] += 1

							}
						}
					}
				}
			}
			// Concatenate two slices
			histogram.Values = append(histogram.Values, regionHistogram...)
		}
	}

	return histogram, nil
}

// checkHistograms check if the histograms are correct.
func checkHistograms(hist1, hist2 []uint64) error {
	if len(hist1) == 0 || len(hist2) == 0 {
		return errors.New("Could not compare the histograms. The histogram is empty.")
	}
	if len(hist1) != len(hist2) {
		return errors.New("Could not compare the histograms. The slices have different sizes.")
	}
	return nil
}

// EuclidDistance returns the euclidean distance between two histograms
func (h *Histogram) EuclidDistance(histogram Histogram) (float64, error) {
	if err := checkHistograms(h.Values, histogram.Values); err != nil {
		return 0.0, err
	}

	var sum float64
	for index := 0; index < len(h.Values); index++ {
		sum += math.Pow(float64(h.Values[index])-float64(histogram.Values[index]), 2)
	}
	return math.Sqrt(sum), nil
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
