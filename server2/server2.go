package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
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
	lissajous(w, r.URL.Query())
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Counter: %d\n", count)
	mu.Unlock()
}

var palette = []color.Color{
	color.RGBA{0x0e, 0x0e, 0x0e, 0xff},
	color.RGBA{0xdd, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xdd, 0xff},
	color.RGBA{0x00, 0xdd, 0x00, 0xff}}

func tryGetQueryAtoi(query url.Values, key string, defaultValue int) int {
	if key == "" {
		return defaultValue
	} else if value := query.Get(key); value != "" {
		if i, err := strconv.Atoi(value); err != nil {
			log.Print(err)
			return defaultValue
		} else {
			return i
		}
	} else {
		return defaultValue
	}
}

func lissajous(out io.Writer, query url.Values) {
	var (
		cycles  = tryGetQueryAtoi(query, "cycles", 5)
		res     = 0.001
		size    = tryGetQueryAtoi(query, "size", 360)
		nframes = tryGetQueryAtoi(query, "nframes", 64)
		delay   = tryGetQueryAtoi(query, "delay", 8)
	)
	freq := rand.Float64() * 6.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(t)%4)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
