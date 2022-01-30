package main

import (
	"log"

	"adaptive-mesh-refinement-art/modules/amr"
	"adaptive-mesh-refinement-art/modules/image"
)

func main() {
	path := "th.jpg"
	img := image.NewImageFromFile(path, 12)
	img.Save("preprocessed.png")

	tree := amr.NewQuadtreeFromImage(img)
	tree.Refine()
	image.NewFromBoolMat(tree.GetLayerMat(3)).Save("layer.png")

	mat := tree.GetAMRMat()
	image.NewFromBoolMat(mat).Save("amr.png")

	log.Println("done")
}
