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

func NewImage(mat *gocv.Mat) *Image {
	size := mat.Size()
	if size[0] != size[1] {
		log.Panicf("the image should be a square: %d != %d", size[0], size[1])
	}

	return &Image{
		mat:  mat,
		size: size[0],
	}
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
	log.Printf("saving %s\n", path)

	if path[len(path)-4:] != ".png" {
		gocv.IMWrite(path, *img.mat)
	}

	params := []int{gocv.IMWritePngCompression, 8}
	gocv.IMWriteWithParams(path, *img.mat, params)
}

func (img *Image) GetMat() [][]bool {
	data, err := img.mat.DataPtrUint8()
	if err != nil {
		log.Panicf("img.mat.DataPrtUint8 failed: %v", err)
	}

	res := make([][]bool, img.size)
	idx := 0
	for i := 0; i < img.size; i++ {
		r := make([]bool, img.size)
		for j := 0; j < img.size; j++ {
			r[j] = data[idx] > 0
			idx++
		}
		res[i] = r
	}
	return res
}

func (img *Image) GetMaxLv() int {
	return int(math.Log2(float64(img.size)))
}
