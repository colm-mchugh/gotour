package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x + y*x), uint8(x + y*y), 255, 255}
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w-1, i.h-1)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{1000, 1000}
	pic.ShowImage(m)
}
