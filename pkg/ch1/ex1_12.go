package ch1

import (
	"fmt"
	"image"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func Ex1_12() {
	http.HandleFunc("/lissajous", lissaJous2Handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissaJous2Handler(writer http.ResponseWriter, request *http.Request) {
	lissaJous2(writer, request)
}

func lissaJous2(out http.ResponseWriter, request *http.Request) {
	var revolutions = 5.0
	var resolution = 0.001
	var err error

	switch {
	case request.FormValue("revolutions") != "":
		revolutions, err = strconv.ParseFloat(request.FormValue("revolutions"), 64)
		if err != nil {
			panic(err)
		}
	case request.FormValue("resolution") != "":
		resolution, err = strconv.ParseFloat(request.FormValue("resolution"), 64)
		if err != nil {
			panic(err)
		}
	}

	frequency := rand.Float64() * 3.0
	animation := gif.GIF{
		LoopCount: NumberOfAnimationFrames,
	}

	phase := 0.0
	for i := 0; i < NumberOfAnimationFrames; i++ {
		rectangle := image.Rect(0, 0, 2*ImageCanvasSize+1, 2*ImageCanvasSize+1)
		img := image.NewPaletted(rectangle, palette)
		for t := 0.0; t < revolutions*2*math.Pi; t += resolution {
			x := math.Sin(t)
			y := math.Sin(t*frequency + phase)
			img.SetColorIndex(ImageCanvasSize+int(x*ImageCanvasSize+0.5), ImageCanvasSize+int(y*ImageCanvasSize+0.5), WhiteIndex)

		}
		phase += 0.1
		animation.Delay = append(animation.Delay, DelayBetweenFrames)
		animation.Image = append(animation.Image, img)
	}
	err = gif.EncodeAll(out, &animation)
	if err != nil {
		err := fmt.Errorf("error during gif.Encode: %w", err)
		fmt.Println(err)
	}

}
