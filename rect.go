package main

import (
	"image"

	"math/rand"
)

type rect struct {
	centerX, centerY, radX, radY int
}

func (s *rect) Rasterize(bounds image.Rectangle) []Line {
	lines := make([]Line, s.radY*2+1)
	for i, y := 0, maxInt(s.centerY-s.radY, 0); y < minInt(s.centerY+s.radY+1, bounds.Max.Y); i, y = i+1, y+1 {
		lines[i].y = y
		lines[i].xMin = maxInt(s.centerX-s.radX, 0)
		lines[i].xMax = minInt(s.centerX+s.radX+1, bounds.Max.X)
	}
	return lines
}

func (s *rect) Copy() Shape {
	return &rect{s.centerX, s.centerY, s.radX, s.radY}
}

func (s *rect) Mutate(bounds image.Rectangle) {
	offset := 20
	radOffset := 10
	s.centerX += rand.Intn(2*offset+1) - offset
	s.centerY += rand.Intn(2*offset+1) - offset
	s.radX += rand.Intn(2*radOffset+1) - radOffset
	s.radY += rand.Intn(2*radOffset+1) - radOffset
	s.centerX = clampInt(s.centerX, bounds.Min.X, bounds.Max.X)
	s.centerY = clampInt(s.centerY, bounds.Min.Y, bounds.Max.Y)
	s.radX = clampInt(s.radX, 2, bounds.Size().X/2)
	s.radY = clampInt(s.radY, 2, bounds.Size().Y/2)
}
