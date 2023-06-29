/*
*

	@author: OneSpanner
	@since: 2023/6/28
	@desc: //TODO

*
*/
package ch1

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// var palette = []color.Color{color.White, color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}
var paletteV2 = []color.Color{
	color.White,
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},
	color.RGBA{0xA9, 0x56, 0x00, 0xFF},
	color.RGBA{0x56, 0xA9, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF},
	color.RGBA{0x56, 0x00, 0xA9, 0xFF},
	color.RGBA{0xA9, 0x00, 0x56, 0xFF},
}

//const (
//	whiteIndex = 0 // first color in palette
//	blackIndex = 1 // next color in palette
//	greenIndex = 2
//)

// Exer6
//
//	 @Description: 练习1.6
//	 @运行命令：
//		 go build main.go
//		 ./main >out.gif
func Exer6() {
	//fmt.Println("\n=====================Exer6 Start=====================")

	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	lissajousV2(os.Stdout)

	//fmt.Println("=====================Exer6 End=====================")
}

func lissajousV2(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, paletteV2)
		colorIndex := uint8(0)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				colorIndex)
			colorIndex = (colorIndex + 1) % 7
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
