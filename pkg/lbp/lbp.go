// Package lbp features all of the functions
// needed in order to generate local binary patterns based
// off a given image.
package lbp

import (
	"image"
)

// GenerateLocalBinaryPatterns is a function that takes in a source
// image and then computes the Local Binary Patterns or LBP for a given
// image.
func GenerateLocalBinaryPattern(src image.Image) error {
	return nil
}

// GetImageSize takes in a source image and returns the
// dimensions of said image
func GetImageSize(src image.Image) (int, int) {
	if src == nil {
		return 0, 0
	}

	// Get the image bounds
	bounds := src.Bounds()
	// Return the width and height
	return bounds.Max.X, bounds.Max.Y
}

// GetPixels is a function that takes in an image and converts
// the image from a 2D array of pixels into a 2D array of uint8
func GetPixels(src image.Image) [][]uint8 {
	var pixels [][]uint8
	width, height := GetImageSize(src)
	for x := 0; x < width; x++ {
		var row []uint8
		for y := 0; y < height; y++ {
			r, g, b, _ := src.At(x, y).RGBA()
			pixel := (float32(r) * 0.299) + (float32(g) * 0.587) + (float32(b) * 0.114)
			row = append(row, uint8(pixel))
		}
		pixels = append(pixels, row)
	}
	return pixels
}
