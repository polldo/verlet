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
	grid    *verlet.Grid
	flagMov = 0

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
	grid.Extract(0, 0).Fixed = true
	grid.Extract(0, grid.Rows-1).Fixed = true
}

func update() {
	grid.Update(6)
}

func pressed(click *pixel.Vec) {
	if click != nil {
		v := verlet.Vector(*click)

		if flagMov == 0 {
			grid.Origin.Position = &v

		} else if flagMov == 1 {
			edge := grid.Extract(grid.Cols-1, 0)
			edge.Fixed = true
			edge.Position = &v

		} else {
			up := grid.Extract(0, 0)
			down := grid.Extract(0, grid.Rows-1)
			dist := up.Position.Sub(down.Position)
			up.Position = &v
			down.Position = up.Position.Sub(dist)
		}
	}
}

func release() {
	if flagMov == 1 {
		edge := grid.Extract(grid.Cols-1, 0)
		edge.Fixed = false
	}
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

		if win.Pressed(pixelgl.MouseButtonLeft) {
			c := win.MousePosition()
			pressed(&c)
		}
		if win.JustReleased(pixelgl.MouseButtonLeft) {
			release()
		}

		update()
		draw(imd)

		imd.Draw(win)
		win.Update()
	}
}
