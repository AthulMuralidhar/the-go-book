package ch3

import (
	"fmt"
	"math"
)

const cells = 100
const width = 600
const height = 320
const xyScale = width / 2 / xyRange
const xyRange = 30.0
const zScale = height * 0.4

var sin30, cos30 = math.Sin(math.Pi / 6), math.Cos(math.Pi / 6)

func Surface() {

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg'  style='stroke: grey;fill: white;stroke-width: 0.7' width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			aX, aY := corner(i+1, j)
			bX, bY := corner(i, j)
			cX, cY := corner(i, j+1)
			dX, dY := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g, %g, %g, %g, %g, %g, %g, %g '/>\n", aX, aY, bX, bY, cX, cY, dX, dY)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)

	z := computeZ(x, y)

	xProjection := width/2 + (x-y)*cos30*xyScale
	yProjection := height/2 + (x+y)*sin30*xyScale - z*zScale

	return xProjection, yProjection
}

func computeZ(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
