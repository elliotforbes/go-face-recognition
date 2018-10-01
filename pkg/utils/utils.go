package utils

import (
	"image"
	_ "image/jpeg"
	"os"
)

// LoadImage takes in a filepath and returns an image.Image or an error
func LoadImage(filePath string) (image.Image, error) {
	fImage, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fImage.Close()

	img, _, err := image.Decode(fImage)
	if err != nil {
		return nil, err
	}

	return img, nil
}
