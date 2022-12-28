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

var changingColor = color.RGBA{R: 0, G: 0, B: 0, A: 0}
var changingPallet = []color.Color{color.Black, changingColor}

// The colours are beautifully switched i.e gif is in dark mode :)

func EX1_6(out io.Writer) {
	frequency := rand.Float64() * 3.0
	animation := gif.GIF{
		LoopCount: NumberOfAnimationFrames,
	}

	phase := 0.0
	for i := 0; i < NumberOfAnimationFrames; i++ {
		rectangle := image.Rect(0, 0, 2*ImageCanvasSize+1, 2*ImageCanvasSize+1)
		img := image.NewPaletted(rectangle, changingPallet)
		for t := 0.0; t < OscillatorRevolutions*2*math.Pi; t += AngularResolution {
			x := math.Sin(t)
			y := math.Sin(t*frequency + phase)
			img.SetColorIndex(ImageCanvasSize+int(x*ImageCanvasSize+0.5), ImageCanvasSize+int(y*ImageCanvasSize+0.5), WhiteIndex)

		}
		phase += 0.1
		animation.Delay = append(animation.Delay, DelayBetweenFrames)
		animation.Image = append(animation.Image, img)
		//changingColor.R += int(rand.Uint32() * 255)
	}
	err := gif.EncodeAll(out, &animation)
	if err != nil {
		err := fmt.Errorf("error during gif.Encode: %w", err)
		fmt.Println(err)
	}
}
