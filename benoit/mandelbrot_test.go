package benoit_test

import (
	"image"
	"testing"

	"github.com/shogg/mandelbrot/benoit"
)

var (
	img = image.NewRGBA(image.Rect(0, 0, 1024, 768))
)

func _BenchmarkMandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benoit.Mandelbrot(-.80, 2.5, img)
	}
}

func BenchmarkMandelbrotParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benoit.MandelbrotParallel(-.80, 2.5, img)
	}
}

func _BenchmarkMandelbrotSampled(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benoit.MandelbrotSampled(-.80, 2.5, img)
	}
}
