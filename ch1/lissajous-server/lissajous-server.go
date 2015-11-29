package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	cycles  = 5.0   // default number of complete x oscillator revolutions
	size    = 250   // default image canvas covers [-size..+size]
	res     = 0.001 // default angular resolution
	nFrames = 64    // default number of animation frames
	delay   = 8     // default delay between frames in 10ms units
)

func main() {
	log.Print("Server running...")
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	if x, err := strconv.ParseFloat(r.FormValue("cycles"), 64); err == nil {
		cycles = x
	}
	if x, err := strconv.Atoi(r.FormValue("size")); err == nil {
		size = x
	}
	if x, err := strconv.ParseFloat(r.FormValue("res"), 64); err == nil {
		res = x
	}
	if x, err := strconv.Atoi(r.FormValue("nframes")); err == nil {
		nFrames = x
	}
	if x, err := strconv.Atoi(r.FormValue("delay")); err == nil {
		delay = x
	}
	log.Printf("cycles=%v\n", cycles)
	log.Printf("size=%v\n", size)
	log.Printf("res=%v\n", res)
	log.Printf("nFrames=%v\n", nFrames)
	log.Printf("delay=%v\n", delay)
	lissajous(w)
}

func lissajous(out io.Writer) {
	palette := []color.Color{color.Black}
	for i := 0; i < nFrames; i++ {
		red := (uint8)((i * 255) / nFrames)
		green := (uint8)(((i + nFrames/2) * 255) / nFrames)
		blue := (uint8)(((nFrames/2 - i) * 255) / nFrames)
		palette = append(palette, color.RGBA{red, green, blue, 0xff})
	}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0
	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			xcoord := size + int(x*(float64)(size)+0.5)
			ycoord := size + int(y*(float64)(size)+0.5)
			img.SetColorIndex(xcoord, ycoord, (uint8)(i+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
