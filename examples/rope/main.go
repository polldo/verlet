package main

import (
	"verlet"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	rope *verlet.Rope

	windowWidth  = 1280.
	windowHeight = 600.
)

func setup() {
	p := &verlet.VerletParams{
		Gravity:  verlet.Vector{X: 0, Y: -0.8},
		Bound:    verlet.Vector{X: windowWidth, Y: windowHeight},
		Friction: 0.9,
	}
	rope = verlet.NewRope(20, 20, p)
}

func update(click *pixel.Vec) {
	if click != nil {
		v := verlet.Vector(*click)
		rope.Head.Position = &v
	}
	rope.Update(60)
}

func draw(imd *imdraw.IMDraw) {
	// for _, p := range rope.Points {
	// 	imd.Color = colornames.Blue
	// 	imd.Push(pixel.V(p.Position.X, p.Position.Y))
	// 	imd.Circle(p.Radius, 0)
	// }

	imd.Color = colornames.Orange
	imd.Push(pixel.V(rope.Head.Position.X, rope.Head.Position.Y))
	imd.Circle(rope.Head.Radius, 0)

	for _, l := range rope.Lines {
		imd.Color = colornames.Orange
		imd.Push(pixel.Vec(*l.A.Position), pixel.Vec(*l.B.Position))
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
		draw(imd)

		imd.Draw(win)
		win.Update()
	}
}
