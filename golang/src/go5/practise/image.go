package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
)

type Image struct {
	Height, Width int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Height, m.Width)
}

func (m Image) At(x, y int) color.Color {
	c := uint8((x + y) / 2)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	m := Image{256, 256}
	ShowImage(m)
}

func ShowImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./src/go5/bluescale.png", buf.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
}
