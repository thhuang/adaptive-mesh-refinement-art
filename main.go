package main

import (
	"image"
	"log"

	"gocv.io/x/gocv" // Install OpenCV first: https://gocv.io/getting-started/macos/
)

func main() {
	path := "th.jpg"
	originalMat := gocv.IMRead(path, gocv.IMReadAnyColor)
	if originalMat.Empty() {
		log.Panic("cannot read image: " + path)
	}

	gaussianBlurMat := gocv.NewMat()
	gocv.GaussianBlur(originalMat, &gaussianBlurMat, image.Point{X: 3, Y: 3}, 0, 0, gocv.BorderDefault)

	grayMat := gocv.NewMat()
	gocv.CvtColor(gaussianBlurMat, &grayMat, gocv.ColorBGRToGray)

	windowForGray := gocv.NewWindow("gray image")
	defer windowForGray.Close()
	windowForGray.IMShow(grayMat)

	cannyMat := gocv.NewMat()
	gocv.Canny(grayMat, &cannyMat, 150, 200)
	gocv.ConvertScaleAbs(cannyMat, &cannyMat, 1, 0)

	windowForCanny := gocv.NewWindow("canny image")
	defer windowForCanny.Close()
	windowForCanny.IMShow(cannyMat)

	// wait
	log.Println("press ESC to terminate")
	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("terminate")
			break
		}
	}

	log.Println("done")
}
