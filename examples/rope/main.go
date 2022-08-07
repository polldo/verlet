package main

import (
	"github.com/polldo/verlet"
	"github.com/polldo/verlet/component"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	windowWidth  = 1280.
	windowHeight = 600.
)

func newRope() *component.Rope {
	return component.NewRope(
		20,
		20,
		verlet.Gravity(0, -0.8),
		verlet.Bound(windowWidth, windowHeight),
		verlet.Friction(0.9),
	)
}

func update(click *pixel.Vec, rope *component.Rope) {
	if click != nil {
		v := verlet.Vector(*click)
		rope.Head.Position = v
	}
	rope.Update(60)
}

func draw(imd *imdraw.IMDraw, rope *component.Rope) {
	imd.Color = colornames.Orange
	imd.Push(pixel.V(rope.Head.Position.X, rope.Head.Position.Y))
	imd.Circle(rope.Head.Radius, 0)

	for _, l := range rope.Lines {
		imd.Color = colornames.Orange
		imd.Push(pixel.Vec(l.A.Position), pixel.Vec(l.B.Position))
		imd.Line(5)
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

	rope := newRope()

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		imd.Clear()

		var click *pixel.Vec = nil
		if win.Pressed(pixelgl.MouseButtonLeft) {
			c := win.MousePosition()
			click = &c
		}

		update(click, rope)
		draw(imd, rope)

		imd.Draw(win)
		win.Update()
	}
}
