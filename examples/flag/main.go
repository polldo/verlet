package main

import (
	"image/color"
	"verlet"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	grid *verlet.Grid

	windowWidth  = 800.
	windowHeight = 300.
)

func setup() {
	p := &verlet.VerletParams{
		Gravity:  verlet.Vector{X: 0.1, Y: -0.1},
		Bound:    verlet.Vector{X: windowWidth, Y: windowHeight},
		Friction: 1,
	}
	cols, rows := 3*6, 12
	dist := 12.
	grid = verlet.NewGrid(cols, rows, dist, p)

	grid.Extract(0, grid.Rows-1).Fixed = true
}

func update(click *pixel.Vec) {
	if click != nil {
		v := verlet.Vector(*click)
		grid.Head.Position = &v
	}

	grid.Update(6)
}

func draw(imd *imdraw.IMDraw) {
	for i, l := range grid.Lines {
		imd.Color = flagColor(i, grid.Rows, grid.Cols)
		imd.Push(pixel.Vec(*l.A.Position), pixel.Vec(*l.B.Position))
		imd.Line(2)
	}
}

func flagColor(i, rows, cols int) color.RGBA {
	if i < rows-1+(rows*2-1)*(cols/3) {
		return colornames.Green
	} else if i < rows-1+(rows*2-1)*(cols/3*2) {
		return colornames.White
	} else {
		return colornames.Red
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
