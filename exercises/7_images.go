package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width  int
	height int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x * y), uint8((x + y) / 2), uint8(x * y), 200}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
