package entity

import (
	"log"

	"gocv.io/x/gocv"
)

func New(mat *gocv.Mat) *Image {
	return &Image{
		mat: mat,
	}
}

type Image struct {
	mat *gocv.Mat
}

func (img *Image) Plot() {
	window := gocv.NewWindow("image")
	defer window.Close()
	window.IMShow(*img.mat)

	log.Println("press ESC to terminate")
	for {
		key := gocv.WaitKey(3)
		if key == 27 {
			log.Println("terminate")
			break
		}
	}
}
