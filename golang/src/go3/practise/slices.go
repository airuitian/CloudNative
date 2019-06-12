package main

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
)

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := range p {
		p[i] = make([]uint8, dx)
	}
	for y, row := range p {
		for x := range row {
			row[x] = uint8(x * y)
		}
	}
	return p
}

func main() {
	Show(Pic)
}

func Show(f func(int, int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}

func ShowImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./src/go3/bluescale.png", buf.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
}
