package ch1

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

const OscillatorRevolutions = 5
const AngularResolution = 0.001
const ImageCanvasSize = 100
const NumberOfAnimationFrames = 64

const BlackIndex = 0
const WhiteIndex = 1
const DelayBetweenFrames = 8

var palette = []color.Color{color.Black, color.White}

func LissaJous1(out io.Writer) {
	frequency := rand.Float64() * 3.0
	animation := gif.GIF{
		LoopCount: NumberOfAnimationFrames,
	}

	phase := 0.0
	for i := 0; i < NumberOfAnimationFrames; i++ {
		rectangle := image.Rect(0, 0, 2*ImageCanvasSize+1, 2*ImageCanvasSize+1)
		img := image.NewPaletted(rectangle, palette)
		for t := 0.0; t < OscillatorRevolutions*2*math.Pi; t += AngularResolution {
			x := math.Sin(t)
			y := math.Sin(t*frequency + phase)
			img.SetColorIndex(ImageCanvasSize+int(x*ImageCanvasSize+0.5), ImageCanvasSize+int(y*ImageCanvasSize+0.5), WhiteIndex)

		}
		phase += 0.1
		animation.Delay = append(animation.Delay, DelayBetweenFrames)
		animation.Image = append(animation.Image, img)
	}
	err := gif.EncodeAll(out, &animation)
	if err != nil {
		err := fmt.Errorf("error during gif.Encode: %w", err)
		fmt.Println(err)
	}
}
