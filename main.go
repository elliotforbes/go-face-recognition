package main

import (
	"fmt"

	"github.com/elliotforbes/go-face-recognition/pkg/utils"
)

func main() {
	fmt.Println("Go Face Recognition Library")

	values := []float64{1.3, 2.3, 4.2, 5.3, 1.1, 1.7}

	histogram := utils.GenerateHistogram(values)
	fmt.Println(histogram)

	distance := utils.CalculateEuclidDistance(utils.Vector3{2, -1, 0}, utils.Vector3{-2, 2, 0})

	fmt.Println(distance)

}
