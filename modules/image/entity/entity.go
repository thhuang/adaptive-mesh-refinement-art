package entity

import (
	"log"
	"math"

	"gocv.io/x/gocv"
)

type Image struct {
	mat  *gocv.Mat
	size int
}

func (img *Image) Init(mat *gocv.Mat) *Image {
	size := mat.Size()
	if size[0] != size[1] {
		log.Panicf("the image should be a square: %d != %d", size[0], size[1])
	}

	img = &Image{
		mat:  mat,
		size: size[0],
	}
	return img
}

func (img *Image) Plot() {
	window := gocv.NewWindow("image")
	defer window.Close()
	window.IMShow(*img.mat)

	log.Println("press ESC to terminate")
	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("terminated")
			break
		}
	}
}

func (img *Image) Save(path string) {
	gocv.IMWrite(path, *img.mat)
}

func (img *Image) GetMat() [][]bool {
	res := make([][]bool, img.size)
	for i := 0; i < img.size; i++ {
		r := make([]bool, img.size)
		for j := 0; j < img.size; j++ {
			r[j] = img.mat.GetUCharAt(i, j) > 0
		}
		res[i] = r
	}
	return res
}

func (img *Image) GetMaxLv() int {
	return int(math.Log2(float64(img.size)))
}
