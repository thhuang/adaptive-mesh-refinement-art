package main

import (
	"log"

	"adaptive-mesh-refinement-art/modules/image"
)

func main() {
	path := "th.jpg"
	img := image.NewImageFromFile(path, 512)

	img.Plot()

	log.Println("done")
}
