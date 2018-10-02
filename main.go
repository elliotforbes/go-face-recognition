package main

import (
	"fmt"

	"github.com/elliotforbes/go-face-recognition/pkg/histogram"
	"github.com/elliotforbes/go-face-recognition/pkg/lbp"
	"github.com/elliotforbes/go-face-recognition/pkg/utils"
)

const testdir = "testdata"

func getHistogram(pathname string) (histogram.Histogram, error) {
	myImage, err := utils.LoadImage(pathname)
	if err != nil {
		fmt.Println(err)
	}

	lbpPixels, err := lbp.LocalBinaryPattern(myImage, 8, 1)
	if err != nil {
		fmt.Println(err)
	}

	histogram, err := histogram.GenerateHistogram(lbpPixels, 8, 8)
	if err != nil {
		fmt.Println(err)
	}

	return histogram, nil
}

func main() {
	fmt.Println("Go Face Recognition Library")

	histogram, err := getHistogram("testdata/avengers-02.jpeg")
	if err != nil {
		fmt.Println(err)
	}

	histogram2, err := getHistogram("testdata/avengers-02.jpeg")
	if err != nil {
		fmt.Println(err)
	}

	distance, err := histogram.EuclidDistance(histogram2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Euclidean Distance: %f\n", distance)

	// err = lbp.GenerateLBPImage(myImage, "out.jpg")
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
