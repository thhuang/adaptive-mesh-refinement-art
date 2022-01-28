package main

import (
	"log"

    "adaptive-mesh-refinement-art/image"
)

func main() {
	path := "th.jpg"
	image.New(path, 1024)

	log.Println("done")
}
