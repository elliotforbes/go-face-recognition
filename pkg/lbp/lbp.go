// Package lbp features all of the functions
// needed in order to generate local binary patterns based
// off a given image.
package lbp

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"strconv"
)

// getBinaryString function used to get a binary value as a string based on a threshold.
// Return "1" if the value is equal or higher than the threshold or "0" otherwise.
func getBinaryString(value, threshold int) string {
	if value >= threshold {
		return "1"
	} else {
		return "0"
	}
}

// GenerateLocalBinaryPatterns is a function that takes in a source
// image and then computes the Local Binary Patterns or LBP for a given
// image.
func LocalBinaryPattern(img image.Image, radius, neighbors uint8) ([][]uint64, error) {

	var lbpPixels [][]uint64
	// Check the parameters
	if img == nil {
		return lbpPixels, errors.New("The image passed to the ApplyLBP function is nil")
	}
	if radius <= 0 {
		return lbpPixels, errors.New("Invalid radius parameter passed to the ApplyLBP function")
	}
	if neighbors <= 0 {
		return lbpPixels, errors.New("Invalid neighbors parameter passed to the ApplyLBP function")
	}

	// Get the pixels 'matrix' ([][]uint8)
	pixels := GetPixels(img)

	// Get the image size (width and height)
	width, height := GetImageSize(img)

	// For each pixel in the image
	for x := 1; x < width-1; x++ {
		var currentRow []uint64
		for y := 1; y < height-1; y++ {

			// Get the current pixel as the threshold
			threshold := int(pixels[x][y])

			binaryResult := ""
			// Window based on the radius (3x3)
			for tempX := x - 1; tempX <= x+1; tempX++ {
				for tempY := y - 1; tempY <= y+1; tempY++ {
					// Get the binary for all pixels around the threshold
					if tempX != x || tempY != y {
						binaryResult += getBinaryString(int(pixels[tempX][tempY]), threshold)
					}
				}
			}

			// Convert the binary string to a decimal integer
			dec, err := strconv.ParseUint(binaryResult, 2, 64)
			if err != nil {
				return lbpPixels, errors.New("Error converting binary to uint in the ApplyLBP function")
			} else {
				// Append the decimal do the result slice
				// ParseUint returns a uint64 so we need to convert it to uint8
				currentRow = append(currentRow, uint64(dec))
			}
		}
		// Append the slice to the 'matrix'
		lbpPixels = append(lbpPixels, currentRow)
	}
	return lbpPixels, nil
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

// GenerateLBPImage takes in an image and creates an
func GenerateLBPImage(img image.Image, output string) error {

	lbpPixels, err := LocalBinaryPattern(img, 16, 1)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	newImage := image.NewRGBA(img.Bounds())
	for x, row := range lbpPixels {
		for y, _ := range row {
			value := lbpPixels[x][y]
			c := color.RGBA{
				uint8(value),
				uint8(value),
				uint8(value),
				uint8(value),
			}
			newImage.Set(x, y, c)
		}
	}

	e1 := jpeg.Encode(file, newImage, nil)
	if e1 != nil {
		fmt.Println(e1)
		return e1
	}
	return nil
}
