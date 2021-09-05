package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	worms []*Worm

	windowWidth  = 1280.
	windowHeight = 600.
)

func setup() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 80; i++ {
		worms = append(worms, NewWorm())
	}
}

func update() {
	for _, w := range worms {
		w.Update()
	}
}

func draw(imd *imdraw.IMDraw) {
	for _, w := range worms {
		w.Draw(imd)
	}
}

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "verlet",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	setup()

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		imd.Clear()

		update()
		draw(imd)

		imd.Draw(win)
		win.Update()
	}
}
