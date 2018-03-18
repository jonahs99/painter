package main

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func clampInt(v, min, max int) int {
	if min >= v {
		return min
	}
	if max <= v {
		return max
	}
	return v
}

func clamp(v, min, max float64) float64 {
	if min >= v {
		return min
	}
	if max <= v {
		return max
	}
	return v
}

func loadImage(path string) (image.Image, error) {
	inFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	img, _, err := image.Decode(inFile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func saveImagePNG(img image.Image, path string) error {
	outFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	return err
}
