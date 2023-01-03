package ch3

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

func Mandelbrot() {
	const width, height = 1024, 1024
	const yMin, yMax = -2, 2
	const xMin, xMax = -2, 2

	pngImage := image.NewRGBA(image.Rect(0, 0, width, height))

	for j := 0; j < height; j++ {
		y := float64(j)/height*(yMax-yMin) + yMin
		for i := 0; i < width; i++ {
			x := float64(i)/width*(xMax-xMin) + xMin
			z := complex(x, y)

			pngImage.Set(i, j, generator(z))
		}
	}
	err := png.Encode(os.Stdout, pngImage)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func generator(z complex128) color.Color {
	const iterations, contrast = 200, 15
	var value complex128
	for i := 0; i < iterations; i++ {
		value = value*value + z
		if cmplx.Abs(value) > 2 {
			return color.Gray{Y: uint8(255 - contrast*i)}
		}
	}
	return color.Black
}
