package main

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
)

type Image struct {
	Height, Width int
}

func main() {
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
