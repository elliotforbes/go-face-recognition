package main

import (
	"fmt"

	"github.com/elliotforbes/go-face-recognition/pkg/histogram"
	"github.com/elliotforbes/go-face-recognition/pkg/utils"
)

const testdir = "testdata"

func getHistogram(pathname string) (histogram.Histogram, error) {
	myImage, err := utils.LoadImage(pathname)
	if err != nil {
		fmt.Println(err)
	}

	histogram, err := histogram.GenerateHistogram(myImage)
	if err != nil {
		fmt.Println(err)
	}

	return histogram, nil
}

func main() {
	fmt.Println("Go Face Recognition Library")

	histogram, err := getHistogram("testdata/png/02.png")
	if err != nil {
		fmt.Println(err)
	}
	histogram2, err := getHistogram("testdata/png/01.png")
	if err != nil {
		fmt.Println(err)
	}

	distance, err := histogram.EuclidDistance(histogram2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Euclidean Distance: %f\n", distance)

	// myImage, _ := utils.LoadImage("testdata/test-100x100.jpg")
	// err = lbp.GenerateLBPImage(myImage, "out.jpg")
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
