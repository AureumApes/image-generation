package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

type vector struct {
	x int
	y int
}

func main() {
	pixels := make(map[vector]color.RGBA, 100)
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{10, 10}})
	colors := []color.RGBA{
		{255 / 2, 0, 0, 0xff},
		{0, 255 / 2, 0, 0xff},
		{0, 0, 255 / 2, 0xff},
		{255 / 2, 255 / 2, 0, 0xff},
	}
	for x := 0; x < 6; x++ {
		for y := 0; y < 10; y++ {
			rand.Seed(time.Now().UnixNano())
			c := colors[rand.Intn(len(colors))]
			img.Set(x, y, c)
			pixels[vector{x, y}] = c
		}
	}
	for x := 5; x < 10; x++ {
		for y := 0; y < 10; y++ {
			var c color.RGBA
			c = pixels[vector{-x + 9, y}]
			img.Set(x, y, c)
			pixels[vector{x, y}] = c
		}
	}
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
