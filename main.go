package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("here we go...")

	target, _ := loadImage("examples/elephant.jpg")
	bounds := target.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	model := NewModel(target)
	fmt.Println("error: ", math.Sqrt(float64(model.errorTot/int64(w*h))))

	nShapes := 1500
	for i := 0; i < nShapes; i++ {
		//lastImg = model.img
		HillClimb(model, 1000)
		model.DrawShape()
		model.shape.Rasterize(bounds)
		//fmt.Println(model.shape.Rasterize(bounds))

		model.computerError()
		fmt.Printf("%d/%d\n", i, nShapes)
		fmt.Println("error: ", math.Sqrt(float64(model.errorTot/int64(w*h))))
		saveImagePNG(&model.img, "examples/out-elephant.png")
		/*if model.errorTot > lastErr {
			fmt.Println("Reject", lastErr, model.errorTot)
			saveImagePNG(&model.img, "examples/hm.png")
			model.img = lastImg
		} else {
			lastErr = model.errorTot
		}*/
	}

}
