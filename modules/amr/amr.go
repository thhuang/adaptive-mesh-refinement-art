package amr

import (
	"adaptive-mesh-refinement-art/modules/amr/entity"
	imageEntity "adaptive-mesh-refinement-art/modules/image/entity"
	"log"
	"math"
)

func NewQuadtreeFromImage(image *imageEntity.Image) *entity.Quadtree {
	log.Println("generate boolean matrix")
	mat := image.GetMat()
	maxLv := image.GetMaxLv()
	size := len(mat) - 1
	root := entity.NewNode(
		0, size,
		entity.NewCoord(0, 0),
		entity.NewCoord(size, 0),
		entity.NewCoord(size, size),
		entity.NewCoord(0, size),
		nil,
	)

	log.Println("generating layer maps")
	lvSize := make([]int, maxLv)
	layers := make(entity.Layers, maxLv)
	for lv := 0; lv < maxLv; lv++ {
		count := int(math.Pow(2, float64(lv)))
		lvSize[lv] = (size) / count
		layers[lv] = make([][]bool, count)
		for c := 0; c < count; c++ {
			layers[lv][c] = make([]bool, count)
		}

		for i := 0; i < len(mat); i++ {
			for j := 0; j < len(mat); j++ {
				if !mat[i][j] {
					continue
				}

				// (2, 5)
				// lv  cnt size qi ri qj rj
				// 0 -> 1 -> 8: (0-2, 0-5)
				// 1 -> 2 -> 4: (0-2, 1-1)
				// 2 -> 4 -> 2: (1-0, 2-1)
				// 3 -> 8 -> 1: (2-0, 5-0)

				ri, rj := i%lvSize[lv], j%lvSize[lv]
				if ri*rj == 0 {
					continue
				}

				qi, qj := i/lvSize[lv], j/lvSize[lv]
				layers.SetSubdivideFlag(lv, qi, qj)
			}
		}
	}

	return entity.NewQuadtree(maxLv, mat, layers, root)
}
