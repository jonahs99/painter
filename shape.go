package main

import (
	"image"

	"math/rand"
)

// Line for rasterizing
type Line struct {
	y, xMin, xMax int
}

// Shape can be drawn and mutated
type Shape interface {
	Rasterize(bounds image.Rectangle) []Line
	Mutate(bounds image.Rectangle)
	Copy() Shape
}

// NewRandomShape gives us a starting shape
func NewRandomShape(bounds image.Rectangle) Shape {
	w, h := bounds.Max.X, bounds.Max.Y
	return &ellipse{
		rand.Intn(w), rand.Intn(h), rand.Intn(50) + 5, rand.Intn(50) + 5,
	}
}
