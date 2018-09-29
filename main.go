package main

import (
	"fmt"

	"github.com/elliotforbes/go-face-recognition/pkg/utils"
)

func main() {
	fmt.Println("Go Face Recognition Library")

	buckets := []utils.Bucket{
		{1, 1.49, nil},
		{1.5, 2.49, nil},
		{2.5, 3.49, nil},
		{3.5, 4.49, nil},
	}
	values := []float64{1.3, 2.3, 4.2, 5.3, 1.1, 1.7}

	histogram := utils.GenerateHistogram(buckets, values)
	fmt.Println(histogram)

	pointA := utils.Vector3{2, -1, 0}
	pointB := utils.Vector3{-2, 2, 0}
	distance := utils.CalculateEuclidDistance(pointA, pointB)

	fmt.Println(distance)

}
