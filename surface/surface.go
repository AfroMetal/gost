package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * .4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.5' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			_, _, z := getXYZ(i, j)
			z *= 10
			alpha := math.Min(1.0, math.Abs(z))
			color := uint32(0xff * alpha)
			if z > 0 {
				color <<= 16
			}
			if anyIsNan([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: #%06x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}
func anyIsNan(floats []float64) bool {
	for _, f := range floats {
		if math.IsNaN(f) {
			return true
		}
	}
	return false
}
func corner(i, j int) (float64, float64) {
	x, y, z := getXYZ(i, j)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}
func getXYZ(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - .5)
	y := xyrange * (float64(j)/cells - .5)
	z := f(x, y)

	return x, y, z
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
