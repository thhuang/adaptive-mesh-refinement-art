package main

import (
	"flag"
	"log"

	"adaptive-mesh-refinement-art/modules/amr"
	"adaptive-mesh-refinement-art/modules/image"
)

func main() {
	input := flag.String("i", "input.png", "Input file")
	output := flag.String("o", "output.png", "Output file")
	lv := flag.Int("lv", 12, "Maximum refinement level")
	cannyParamDesc := "Canny finds edges in an image using the Canny algorithm.\nThe function finds edges in the input image image and marks them in the output map edges using the Canny algorithm.\nThe smallest value between t1 and t2 is used for edge linking. The largest value is used to find initial segments of strong edges.\nReferences:\n- https://en.wikipedia.org/wiki/Canny_edge_detector\n- https://docs.opencv.org/4.5.5/da/d22/tutorial_py_canny.html\n"
	t1 := flag.Float64("t1", 50, cannyParamDesc)
	t2 := flag.Float64("t2", 70, cannyParamDesc)
	flag.Parse()

	img := image.NewImageFromFile(*input, *lv, float32(*t1), float32(*t2))

	tree := amr.NewQuadtreeFromImage(img)
	tree.Refine()

	mat := tree.GetAMRMat()
	image.NewFromBoolMat(mat).Save(*output)

	log.Println("done")
}
