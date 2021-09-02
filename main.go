package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	Gravity  = 0.6
	Friction = 0.9
)

var (
	imd    *imdraw.IMDraw
	points []*Point
	lines  []*Line

	windowWidth  = 1280.
	windowHeight = 300.
)

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

	imd = imdraw.New(nil)

	setup()

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		imd.Clear()

		var click *pixel.Vec = nil
		if win.Pressed(pixelgl.MouseButtonLeft) {
			c := win.MousePosition()
			click = &c
		}

		update(click)
		imd.Draw(win)
		win.Update()
	}
}

func setup() {
	points = append(points, NewPoint(20, 20, true))
	points = append(points, NewPoint(55, 55, false))
	points = append(points, NewPoint(10, 60, false))
	points = append(points, NewPoint(80, 40, false))

	lines = append(lines, NewLine(points[0], points[1]))
	lines = append(lines, NewLine(points[2], points[3]))
	lines = append(lines, NewLine(points[1], points[2]))
}

func update(click *pixel.Vec) {
	if click != nil {
		points[0].Position = *click
	}
	for _, p := range points {
		p.Update()
	}
	for _, l := range lines {
		l.Update()
	}

	for _, l := range lines {
		l.Draw(imd)
	}
	for _, p := range points {
		p.Draw(imd)
	}
}
