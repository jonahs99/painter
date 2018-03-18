package main

import (
	"image"
	"math"

	"math/rand"
)

type ellipse struct {
	centerX, centerY, radX, radY int
}

func (s *ellipse) Rasterize(bounds image.Rectangle) []Line {
	lines := make([]Line, s.radY*2+1)
	//fmt.Println("rasterizing")
	for i, y := 0, maxInt(s.centerY-s.radY, 0); y < minInt(s.centerY+s.radY+1, bounds.Max.Y); i, y = i+1, y+1 {
		lines[i].y = y
		// We solve the equaiton for the ellipse
		// x^2/a^2 + y^2/b^2 = 1 (solving for x at the given y)
		// x = +/- a*sqrt(1 - y^2/b^2)
		yOrigin := float64(y - s.centerY)
		discr := 1.0 - float64(yOrigin*yOrigin)/float64(s.radY*s.radY)
		if discr > 0 {
			w := int(float64(s.radX) * math.Sqrt(discr))
			//fmt.Println(w)
			lines[i].xMin = maxInt(s.centerX-w, 0)
			lines[i].xMax = minInt(s.centerX+w+1, bounds.Max.X)
		} else {
			lines[i].xMax = -1
		}
	}
	//fmt.Println(lines)
	return lines
}

func (s *ellipse) Copy() Shape {
	return &ellipse{s.centerX, s.centerY, s.radX, s.radY}
}

func (s *ellipse) Mutate(bounds image.Rectangle) {
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

	if s.radX*5 < s.radY {
		s.radX = s.radY / 5
	}
	if s.radY*5 < s.radX {
		s.radY = s.radX / 5
	}
}
