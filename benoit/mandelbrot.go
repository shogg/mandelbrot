package benoit

import (
	"image"
	"image/color"
	"math"
	"sync"
)

var (
	maxit = 100
	white = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	black = color.RGBA{R: 0, G: 15, B: 32, A: 255}
)

// Mandelbrot computes an image around c with radius r.
func Mandelbrot(c complex128, r float64, img *image.RGBA) {

	max := img.Bounds().Max
	scale := r / math.Min(float64(max.X), float64(max.Y))

	c0 := c - complex(r/2, r/2)
	for y := 0; y < max.Y; y++ {
		cy := c0 + complex(0, (float64(y)*scale))

		for x := 0; x < max.X; x++ {
			cx := complex(float64(x)*scale, 0)

			it := Julia(cy+cx, maxit)
			if it < maxit {
				img.Set(x, y, white)
			} else {
				img.Set(x, y, black)
			}
		}
	}
}

// MandelbrotParallel computes an image around c with radius r multiple lines in parallel.
func MandelbrotParallel(c complex128, r float64, img *image.RGBA) {

	max := img.Bounds().Max
	scale := r / math.Min(float64(max.X), float64(max.Y))

	var wg sync.WaitGroup
	wg.Add(max.Y)

	c0 := c - complex(r/2, r/2)
	for y := 0; y < max.Y; y++ {
		cy := c0 + complex(0, float64(y)*scale)

		go func(y int) {
			for x := 0; x < max.X; x++ {
				cx := complex(float64(x)*scale, 0)

				it := Julia(cy+cx, maxit)
				if it < maxit {
					img.Set(x, y, white)
				} else {
					img.Set(x, y, black)
				}
			}
			wg.Done()
		}(y)
	}

	wg.Wait()
}

// MandelbrotSampled computes an image around c with radius r each pixel subsampled.
func MandelbrotSampled(c complex128, r float64, img *image.RGBA) {

	const samples = 2
	var sampleColors [samples * samples]color.RGBA

	max := img.Bounds().Max
	scale := r / math.Min(float64(max.X), float64(max.Y))
	subscale := scale / samples

	c0 := c - complex(r/2, r/2)
	for y := 0; y < max.Y; y++ {
		for x := 0; x < max.X; x++ {

			for sy := 0; sy < samples; sy++ {
				cy := c0 + complex(0, float64(y)*scale+float64(sy)*subscale)

				for sx := 0; sx < samples; sx++ {
					cx := complex(float64(x)*scale+float64(sx)*subscale, 0)

					it := Julia(cy+cx, maxit)
					if it < maxit {
						sampleColors[sy*samples+sx] = white
					} else {
						sampleColors[sy*samples+sx] = black
					}
				}
			}

			var r, g, b int
			for _, sc := range sampleColors {
				r += int(sc.R)
				g += int(sc.G)
				b += int(sc.B)
			}
			blend := color.RGBA{
				uint8(r / samples / samples),
				uint8(g / samples / samples),
				uint8(b / samples / samples),
				255}
			img.Set(x, y, blend)
		}
	}
}
