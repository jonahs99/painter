package main

import (
	"image"

	"math/rand"
)

type square struct {
	centerX, centerY, rad int
}

func (s *square) Rasterize(bounds image.Rectangle) []Line {
	lines := make([]Line, s.rad*2+1)
	for i, y := 0, maxInt(s.centerY-s.rad, 0); y < minInt(s.centerY+s.rad+1, bounds.Max.Y); i, y = i+1, y+1 {
		lines[i].y = y
		lines[i].xMin = maxInt(s.centerX-s.rad, 0)
		lines[i].xMax = minInt(s.centerX+s.rad+1, bounds.Max.X)
	}
	return lines
}

func (s *square) Copy() Shape {
	return &square{s.centerX, s.centerY, s.rad}
}

func (s *square) Mutate(bounds image.Rectangle) {
	offset := 20
	radOffset := 10
	s.centerX += rand.Intn(2*offset+1) - offset
	s.centerY += rand.Intn(2*offset+1) - offset
	s.rad += rand.Intn(2*radOffset+1) - radOffset
	s.centerX = clampInt(s.centerX, bounds.Min.X, bounds.Max.X)
	s.centerY = clampInt(s.centerY, bounds.Min.Y, bounds.Max.Y)
	s.rad = clampInt(s.rad, 2, 100)
}
