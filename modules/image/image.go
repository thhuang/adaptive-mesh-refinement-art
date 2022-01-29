package image

import (
	"image"
	"log"

	"gocv.io/x/gocv" // Install OpenCV first: https://gocv.io/getting-started/macos/

	"adaptive-mesh-refinement-art/modules/image/entity"
)

func NewImageFromFile(path string, size int) *entity.Image {
	originalMat := gocv.IMRead(path, gocv.IMReadAnyColor)
	if originalMat.Empty() {
		log.Panic("cannot read image: " + path)
	}

	gaussianBlurMat := gaussianBlur(&originalMat)

	grayMat := gray(gaussianBlurMat)
	// windowForGray := gocv.NewWindow("gray image")
	// defer windowForGray.Close()
	// windowForGray.IMShow(*grayMat)

	cropMat := crop(grayMat)
	// windowForCrop := gocv.NewWindow("cropped image")
	// defer windowForCrop.Close()
	// windowForCrop.IMShow(*cropMat)

	resizedMat := resize(cropMat, size)
	// windowForResized := gocv.NewWindow("resized image")
	// defer windowForResized.Close()
	// windowForResized.IMShow(*resizedMat)

	cannyMat := canny(resizedMat)
	// windowForCanny := gocv.NewWindow("canny image")
	// defer windowForCanny.Close()
	// windowForCanny.IMShow(*cannyMat)

	return entity.New(cannyMat)
}

func gaussianBlur(mat *gocv.Mat) *gocv.Mat {
	res := gocv.NewMat()
	gocv.GaussianBlur(*mat, &res, image.Point{X: 3, Y: 3}, 0, 0, gocv.BorderDefault)
	return &res
}

func gray(mat *gocv.Mat) *gocv.Mat {
	res := gocv.NewMat()
	gocv.CvtColor(*mat, &res, gocv.ColorBGRToGray)
	return &res
}

func crop(mat *gocv.Mat) *gocv.Mat {
	size := mat.Size()
	length := min(size)

	x0 := (size[1] - length) / 2
	y0 := (size[0] - length) / 2
	x1 := size[1] - x0
	y1 := size[0] - y0

	res := mat.Region(image.Rect(x0, y0, x1, y1))
	return &res
}

func resize(mat *gocv.Mat, size int) *gocv.Mat {
	res := gocv.NewMat()
	gocv.Resize(*mat, &res, image.Point{X: size, Y: size}, 0, 0, 1)
	return &res
}

func canny(mat *gocv.Mat) *gocv.Mat {
	res := gocv.NewMat()
	gocv.Canny(*mat, &res, 150, 200)
	gocv.ConvertScaleAbs(res, &res, 1, 0)
	return &res
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
