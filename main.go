package main

import (
	"image"
	"image/png"
	"os"

	"github.com/shogg/mandelbrot/benoit"
)

func main() {

	bounds := image.Rect(0, 0, 2048, 2048)
	img := image.NewRGBA(bounds)

	benoit.MandelbrotSampled(-.80, 2.5, img)

	if err := savePNG("mandel.png", img); err != nil {
		panic(err)
	}
}

func savePNG(name string, img image.Image) error {

	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
