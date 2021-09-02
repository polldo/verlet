package main

import (
	"verlet"

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
	points []*verlet.Point
	lines  []*verlet.Line

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
	p := &verlet.PointParams{0, -0.6, windowWidth, windowHeight, Friction}
	points = append(points, verlet.NewPoint(20, 20, true, colornames.Orange, p))
	points = append(points, verlet.NewPoint(55, 55, false, colornames.Orange, p))
	points = append(points, verlet.NewPoint(10, 60, false, colornames.Orange, p))
	points = append(points, verlet.NewPoint(80, 40, false, colornames.Orange, p))

	lines = append(lines, verlet.NewLine(points[0], points[1], colornames.Black))
	lines = append(lines, verlet.NewLine(points[2], points[3], colornames.Black))
	lines = append(lines, verlet.NewLine(points[1], points[2], colornames.Black))
}

func update(click *pixel.Vec) {
	if click != nil {
		v := verlet.Vector(*click)
		points[0].Position = &v
	}
	for _, p := range points {
		p.Update()
	}
	for _, l := range lines {
		l.Update()
	}

	for _, l := range lines {
		imd.Color = l.Color
		imd.Push(pixel.Vec(*l.A.Position), pixel.Vec(*l.B.Position))
		imd.Line(2)
	}
	for _, p := range points {
		imd.Color = p.Color
		imd.Push(pixel.V(p.Position.X, p.Position.Y))
		imd.Circle(5, 0)
	}
}
