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
	"os"
	"time"
)

const (
	bgIndex = 0     // background color palette index
	fgIndex = 1     // foreground color palette index
	cycles  = 5     // number of complete x oscillator revolutions
	res     = 0.001 // angular resolution
	size    = 100   // image canvas covers [-size..+size]
	nFrames = 64    // number of animation frames
	delay   = 8     // delay between frames in 10ms units
)

func main() {
	log.Print("Server running...")
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}

func lissajous(out io.Writer) {
	palette := []color.Color{color.Black}
	for i := 0; i < nFrames; i++ {
		red := (uint8)((i * 255) / nFrames)
		green := (uint8)(((i + nFrames/2) * 255) / nFrames)
		blue := (uint8)(((nFrames/2 - i) * 255) / nFrames)
		fmt.Fprintf(os.Stderr, "RGB=(0x%02x%02x%02x)\n", red, green, blue)
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
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				(uint8)(i+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
