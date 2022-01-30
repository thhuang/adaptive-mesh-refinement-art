package main

import (
	"log"

	"adaptive-mesh-refinement-art/modules/amr"
	"adaptive-mesh-refinement-art/modules/image"
)

func main() {
	path := "th.jpg"
	img := image.NewImageFromFile(path, 13)
	img.Save("preprocessed.png")

	tree := amr.NewQuadtreeFromImage(img)
	tree.Refine()

	mat := tree.GetAMRMat()
	amrImg := image.NewFromBoolMat(mat)
	amrImg.Save("amr.png")

	log.Println("done")
}
