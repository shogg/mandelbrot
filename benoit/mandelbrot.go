package benoit

import (
	"image"
	"image/color"
	"math"
)

var (
	maxit = 100
	white = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	black = color.RGBA{R: 0, G: 15, B: 32, A: 255}
)

// Mandelbrot computes an image around c with radius r.
// imgres image resulution.
func Mandelbrot(c complex128, r float64, imgres int) image.Image {

	bounds := image.Rect(0, 0, imgres, imgres)
	img := image.NewRGBA(bounds)

	c0 := c - complex(r/2, r/2)
	for y := 0; y < bounds.Max.Y; y++ {
		cy := c0 + complex(0, r*float64(y)/float64(bounds.Max.Y))

		for x := 0; x < bounds.Max.X; x++ {
			cx := complex(r*float64(x)/float64(bounds.Max.X), 0)

			it := Julia(cy+cx, maxit)
			if it < maxit {
				img.Set(x, y, white)
			} else {
				img.Set(x, y, black)
			}
		}
	}

	return img
}

// Julia repeats z = 0; z = z^2 + c until z diverges.
// Stops after max repetitions.
func Julia(c complex128, max int) int {

	var z complex128
	for i := 0; i < max; i++ {
		z = z*z + c

		if mag(z) >= 2 {
			return i
		}
	}

	return max
}

// mag approximates complex number magnitude sqrt(real^2 + imag^2).
func mag(z complex128) float64 {
	return math.Max(math.Abs(real(z)), math.Abs(imag(z)))
}
