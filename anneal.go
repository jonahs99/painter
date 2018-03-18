package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Annealable can be optimized with anneal
type Annealable interface {
	Energy() float64
	DoMove() interface{}
	UndoMove(interface{})
}

// Anneal does simulated annealing!
func Anneal(model Annealable, iterations int) {
	// T = Tmax * (alpha) ^ x
	// Tmin = Tmax * alpha^iterations
	// alpha = Tmin/Tmax ^ 1/its

	Tmax := 1000.0
	Tmin := 0.01
	alpha := math.Pow(Tmin/Tmax, 1.0/float64(iterations))
	fmt.Println(alpha)

	acceptThresh := 6
	maxTries := 200

	T := Tmax
	currentE := model.Energy()
	for i := 0; i < iterations; i++ {
		T *= alpha
		accepted := 0
		tries := 0
		for accepted < acceptThresh {
			revert := model.DoMove()
			newE := model.Energy()
			P := math.Exp(-(newE - currentE) / T)
			if newE < currentE || P > rand.Float64() {
				currentE = newE
				accepted++
			} else {
				model.UndoMove(revert)
			}
			tries++
			if tries > maxTries {
				return
			}
		}
	}
}

// HillClimb does hill climb
func HillClimb(model Annealable, iterations int) {
	currentE := model.Energy()
	for i := 0; i < iterations; i++ {
		revert := model.DoMove()
		newE := model.Energy()
		if newE <= currentE {
			currentE = newE
		} else {
			model.UndoMove(revert)
		}
	}
}
