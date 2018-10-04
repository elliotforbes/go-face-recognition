package histogram_test

import (
	"fmt"
	"testing"

	"github.com/elliotforbes/go-face-recognition/pkg/histogram"
	"github.com/elliotforbes/go-face-recognition/pkg/utils"
)

func TestGenerateHistogram(t *testing.T) {

	image, err := utils.LoadImage("testdata/test-100x100-white.jpg")
	if err != nil {
		t.Error(err)
	}
	histogram, err := histogram.GenerateHistogram(image)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(histogram.Values)

}
