// Lissajous generates GIF animation of random lissajous figures
package main

import (
  "image"
  "image/color"
  "image/gif"
  "io"
  "math"
  "math/rand"
  "os"
)

var palette = []color.Color{color.White, color.Black}

const (
  whiteIndex = 0 // first color in palette
  blackIndex = 1 // next color in palatte
)

func main() {
  lissajous(os.Stdout)
}

func lissajous(out io.Writer){
  const (
    cycles = 5 // Numbers of cycles
    res = 0.001 // Angular resoloution
    size = 100 // Image canvas covers  [-size..+size]
    nframes = 64 // Number of animation frames
    delay = 8 // Delay between frames in 10ms units
  )

  freq := rand.Float64() * 3.0 // Relative frequency of y oscillator
  anim := gif.GIF{LoopCount: nframes}
  phase := 0.0 // Phase difference

  for i := 0 ; i < nframes ; i++ {
    rect := image.Rect(0,0, 2*size+1, 2*size+1)
    img := image.NewPaletted(rect, palette)
    for t := 0.0 ; t < cycles*2*math.Pi; t += res {
      x := math.Sin(t)
      y := math.Sin(t*freq + phase)
      img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),blackIndex)
    }
    phase += 0.1
    anim.Delay = append(anim.Delay,delay)
    anim.Image = append(anim.Image,img)
  }
  gif.EncodeAll(out, &anim)
}
