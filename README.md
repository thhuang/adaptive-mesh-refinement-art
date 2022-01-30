# Adaptive Mesh Refinement Art
Create a cool image with Adaptive Mesh Refinement!

## Dependencies

- [Go](https://go.dev/): download Go version 1.17.6 from [go.dev](https://go.dev/dl/)
- [OpenCV](https://opencv.org/): install with `brew install opencv`

## Usage

- Prepare an image
- Execute with [the script](./scripts/main.go)

  ```go
  package main

  import (
      "log"

      "adaptive-mesh-refinement-art/modules/amr"
      "adaptive-mesh-refinement-art/modules/image"
  )

  func main() {
      path := "<path_to_your_image>"
      img := image.NewImageFromFile(path, 12)
      img.Save("preprocessed.png")

      tree := amr.NewQuadtreeFromImage(img)
      tree.Refine()

      mat := tree.GetAMRMat()
      amrImg := image.NewFromBoolMat(mat)
      amrImg.Save("amr.png")

      log.Println("done")
  }
  ```
