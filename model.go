package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"
)

// Model is the anneable thing
type Model struct {
	target   image.Image
	img      image.RGBA
	errorMap [][]int
	errorTot int64

	shape Shape
	clr   color.Color
}

type modelMoveRevert struct {
	shape Shape
	clr   color.Color
}

// NewModel initializes a Model ready to anneal
func NewModel(target image.Image) *Model {
	blank := *image.NewRGBA(target.Bounds())
	draw.Draw(&blank, blank.Bounds(), image.NewUniform(averageColor(target)), blank.Bounds().Min, draw.Src)

	errorMap := make([][]int, blank.Bounds().Max.Y)
	for i := 0; i < blank.Bounds().Max.Y; i++ {
		errorMap[i] = make([]int, blank.Bounds().Max.X)
	}

	m := Model{
		target:   target,
		img:      blank,
		errorMap: errorMap,
		shape:    NewRandomShape(target.Bounds()),
		clr:      nil,
	}
	m.computerError()

	return &m
}

// Energy scores the current shapes improvement to the target (lower is better)
func (m *Model) Energy() float64 {
	energy, clr := m.computeShape()
	m.clr = clr
	return energy
}

// DoMove produces a random neighbor state for the annealing process
func (m *Model) DoMove() interface{} {
	revert := modelMoveRevert{m.shape.Copy(), m.clr}

	if rand.Intn(10) > 0 {
		m.shape.Mutate(m.target.Bounds())
	} else {
		m.shape = NewRandomShape(m.target.Bounds())
	}

	return revert
}

// UndoMove reverts a previous move in the case that the random neighbor is rejected
func (m *Model) UndoMove(revert interface{}) {
	r := revert.(modelMoveRevert)
	m.shape = r.shape
	m.clr = r.clr
}

// DrawShape draws the current shape onto img
func (m *Model) DrawShape() {
	r, g, b, _ := m.clr.RGBA()

	lines := m.shape.Rasterize(m.img.Bounds())
	for _, line := range lines {
		//fmt.Println(line.y, line.xMin, line.xMax)
		for x := line.xMin; x < line.xMax; x++ {
			or, og, ob, _ := m.img.At(x, line.y).RGBA()
			blendClr := color.RGBA{uint8(((or + r*2) / 3) >> 8), uint8(((og + g*2) / 3) >> 8), uint8(((ob + b*2) / 3) >> 8), 255}
			m.img.Set(x, line.y, blendClr)
		}
	}
	m.computerError()

	// NEXT!
	m.shape = NewRandomShape(m.target.Bounds())
}

func averageColor(img image.Image) color.Color {
	bounds := img.Bounds()

	nPix := 0
	avgR, avgG, avgB := 0, 0, 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			avgR += int(r >> 8)
			avgG += int(g >> 8)
			avgB += int(b >> 8)
			nPix++
		}
	}
	avgR /= nPix
	avgG /= nPix
	avgB /= nPix
	return color.RGBA{uint8(avgR), uint8(avgG), uint8(avgB), 255}
}

func (m *Model) computerError() {
	m.errorTot = 0
	for y := m.img.Bounds().Min.Y; y < m.img.Bounds().Max.Y; y++ {
		for x := m.img.Bounds().Min.X; x < m.img.Bounds().Max.X; x++ {
			ir, ig, ib, _ := m.img.At(x, y).RGBA()
			tr, tg, tb, _ := m.target.At(x, y).RGBA()
			ir8, ig8, ib8 := int(ir>>8), int(ig>>8), int(ib>>8)
			tr8, tg8, tb8 := int(tr>>8), int(tg>>8), int(tb>>8)
			se := (ir8-tr8)*(ir8-tr8) + (ig8-tg8)*(ig8-tg8) + (ib8-tb8)*(ib8-tb8)
			m.errorMap[y][x] = se
			m.errorTot += int64(se)
		}
	}
}

func (m *Model) computeShape() (float64, color.Color) {
	lines := m.shape.Rasterize(m.img.Bounds())
	// First, compute the average color
	nPix := 0
	avgR, avgG, avgB := 0, 0, 0
	for _, line := range lines {
		for x := line.xMin; x < line.xMax; x++ {
			r, g, b, _ := m.target.At(x, line.y).RGBA()
			avgR += int(r >> 8)
			avgG += int(g >> 8)
			avgB += int(b >> 8)
			nPix++
		}
	}
	avgR /= nPix
	avgG /= nPix
	avgB /= nPix
	// Sum the error difference
	squaredError := m.errorTot
	for _, line := range lines {
		for x := line.xMin; x < line.xMax; x++ {
			r, g, b, _ := m.target.At(x, line.y).RGBA()
			r8, g8, b8 := int(r>>8), int(g>>8), int(b>>8)
			variance := (r8-avgR)*(r8-avgR) + (g8-avgG)*(g8-avgG) + (b8-avgB)*(b8-avgB)
			improvement := m.errorMap[line.y][x] - variance
			squaredError -= int64(improvement)
		}
	}

	pix := m.img.Bounds().Size().X * m.img.Bounds().Size().Y
	energy := math.Sqrt(float64(squaredError) / float64(pix))

	return energy, color.RGBA{uint8(avgR), uint8(avgG), uint8(avgB), 255}
}
