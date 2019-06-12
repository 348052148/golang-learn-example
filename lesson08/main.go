package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"sync"
)
var count int
var mu sync.Mutex

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackInde  = 1
)

type myHandler struct{}

func (self myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Write(bytes.NewBufferString("test").Bytes())
}
func main() {
	handler := myHandler{}
	mux := http.NewServeMux()
	mux.Handle("/test", handler)
	http.HandleFunc("/counts", func(w http.ResponseWriter, r *http.Request) {
		buf := bufio.NewWriter(w)
		mu.Lock()
		buf.WriteString(fmt.Sprintf("counst %d", count))
		mu.Unlock()
		buf.Flush()
	})
	http.HandleFunc("/s",func(w http.ResponseWriter, r *http.Request) {
		buf := bufio.NewWriter(w)
		mu.Lock()
		count++
		mu.Unlock()
		buf.WriteString("welcom to "+ r.URL.Path)
		buf.Flush()
	})
	http.HandleFunc("/pic", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})

	http.ListenAndServe(":8080", mux)
}

func lissajous(out io.Writer)  {
	const (
		cycles = 5
		res	= 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount:nframes}
	phase := 0.0
	for i:=0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t:=0.0; t < cycles * 2 * math.Pi; t+=res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size + int(x*size+0.5), size+int(y*size+0.5),blackInde)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}