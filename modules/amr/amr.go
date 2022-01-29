package amr

import (
	"adaptive-mesh-refinement-art/modules/amr/entity"
	imageEntity "adaptive-mesh-refinement-art/modules/image/entity"
)

func NewQuadtreeFromImage(image *imageEntity.Image) *entity.Quadtree {
	return new(entity.Quadtree).Init(image.GetMat(), image.GetMaxLv())
}
