package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/count", counter)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, r.URL.Query())
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Counter: %d\n", count)
	mu.Unlock()
}

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * .4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func surface(out io.Writer, query url.Values) {
	color, err := strconv.ParseUint(query.Get("color"), 16, 32)
	if err != nil {
		log.Print("Wrong color, using #23cc23")
		color = 0x23cc23
	}

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.5' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if anyIsNan([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				continue
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: #%06x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(out, "</svg>")
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
