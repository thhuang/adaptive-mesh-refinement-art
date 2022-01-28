package image

import (
	img "image"
	"log"

	"gocv.io/x/gocv" // Install OpenCV first: https://gocv.io/getting-started/macos/
)

func New(path string, size int) *image {
	originalMat := gocv.IMRead(path, gocv.IMReadAnyColor)
	if originalMat.Empty() {
		log.Panic("cannot read image: " + path)
	}

	gaussianBlurMat := gocv.NewMat()
	gocv.GaussianBlur(originalMat, &gaussianBlurMat, img.Point{X: 3, Y: 3}, 0, 0, gocv.BorderDefault)

	grayMat := gocv.NewMat()
	{
		gocv.CvtColor(gaussianBlurMat, &grayMat, gocv.ColorBGRToGray)

		windowForGray := gocv.NewWindow("gray image")
		defer windowForGray.Close()
		windowForGray.IMShow(grayMat)
	}

	var cropMat gocv.Mat
	{
		size := grayMat.Size()
		length := min(size)

		x0 := (size[1] - length) / 2
		y0 := (size[0] - length) / 2
		x1 := size[1] - x0
		y1 := size[0] - y0

		cropMat = grayMat.Region(img.Rect(x0, y0, x1, y1))

		windowForCrop := gocv.NewWindow("cropped image")
		defer windowForCrop.Close()
		windowForCrop.IMShow(cropMat)
	}

	resizedMat := gocv.NewMat()
	{
		gocv.Resize(cropMat, &resizedMat, img.Point{X: size, Y: size}, 0, 0, 1)

		windowForResized := gocv.NewWindow("resized image")
		defer windowForResized.Close()
		windowForResized.IMShow(resizedMat)
	}

	cannyMat := gocv.NewMat()
	{
		gocv.Canny(resizedMat, &cannyMat, 150, 200)
		gocv.ConvertScaleAbs(cannyMat, &cannyMat, 1, 0)

		windowForCanny := gocv.NewWindow("canny image")
		defer windowForCanny.Close()
		windowForCanny.IMShow(cannyMat)
	}

	log.Println("press ESC to terminate")
	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("terminate")
			break
		}
	}

	return &image{}
}

type image struct {
}

func min(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return res
}
