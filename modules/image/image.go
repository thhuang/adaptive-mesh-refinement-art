package image

import (
	"image"
	"log"
	"math"

	"gocv.io/x/gocv" // Install OpenCV first: https://gocv.io/getting-started/macos/

	"adaptive-mesh-refinement-art/modules/image/entity"
)

func NewFromBoolMat(mat [][]bool) *entity.Image {
	size := len(mat)
	m := gocv.NewMatWithSize(size, size, gocv.MatTypeCV8U)
	for i, r := range mat {
		for j, v := range r {
			if v {
				m.SetUCharAt(i, j, 255)
			}
		}
	}
	return new(entity.Image).Init(&m)
}

func NewImageFromFile(path string, level int) *entity.Image {
	if level > 12 {
		log.Panicf("level cannot be greater than 12: %d > 12", level)
	}
	size := int(math.Pow(2, float64(level))) + 1

	originalMat := gocv.IMRead(path, gocv.IMReadAnyColor)
	if originalMat.Empty() {
		log.Panic("cannot read image: " + path)
	}

	cropMat := crop(&originalMat)
	resizedMat := resize(cropMat, size)
	gaussianBlurMat := gaussianBlur(resizedMat)
	grayMat := gray(gaussianBlurMat)
	cannyMat := canny(grayMat)

	return new(entity.Image).Init(cannyMat)
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
	length := int(math.Min(float64(size[0]), float64(size[1])))

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
	gocv.Canny(*mat, &res, 20, 70)
	gocv.ConvertScaleAbs(res, &res, 1, 0)
	return &res
}
