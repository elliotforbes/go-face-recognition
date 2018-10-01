package main

import (
	"fmt"

	"github.com/elliotforbes/go-face-recognition/pkg/lbp"
	"github.com/elliotforbes/go-face-recognition/pkg/utils"
)

const testdir = "testdata"

func main() {
	fmt.Println("Go Face Recognition Library")
	image, err := utils.LoadImage("testdata/avengers-02.jpeg")
	if err != nil {
		fmt.Println(err)
	}

	pixels := lbp.GetPixels(image)
	fmt.Println(pixels[0][0])
}
