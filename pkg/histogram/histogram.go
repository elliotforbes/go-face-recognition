package histogram

import (
	"errors"
	"fmt"
	"image"
	"math"

	"github.com/elliotforbes/go-face-recognition/pkg/lbp"
)

const GridX = 8
const GridY = 8

type Histogram struct {
	Values []uint64
}

// GenerateHistogram generates a histogram based off the pixel 2D array passed into it
func GenerateHistogram(img image.Image) (Histogram, error) {
	fmt.Println("Generating Histogram")
	var histogram Histogram

	pixels, err := lbp.LocalBinaryPattern(img, 8, 1)
	if err != nil {
		return histogram, errors.New("Error generating LBP")
	}

	height := len(pixels)
	width := len(pixels[0])

	if GridX >= height || GridY >= width {
		return histogram, errors.New("Grid sizes exceed Pixel Dimensions")
	}

	gridWidth := width / int(GridX)
	gridHeight := height / int(GridY)

	fmt.Printf("Grid Width: %v Grid Height: %v\n", gridWidth, gridHeight)

	// Calculates the histogram of each grid
	for gridX := 0; gridX < int(GridX); gridX++ {
		for gridY := 0; gridY < int(GridY); gridY++ {

			// Create a slice with empty 256 positions
			regionHistogram := make([]uint64, 256)

			// Define the start and end positions for the following loop
			startPosX := gridX * gridWidth
			startPosY := gridY * gridHeight
			endPosX := (gridX + 1) * gridWidth
			endPosY := (gridY + 1) * gridHeight

			if gridX == int(GridX)-1 {
				endPosX = width
			}
			if gridY == int(GridY)-1 {
				endPosY = height
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
