package main

import (
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
	p := &verlet.Params{0.1, -0.1, windowWidth, windowHeight, 1}
	cols, rows := 3*6, 12
	dist := 12.
	grid = verlet.NewGrid(cols, rows, dist, p)

	grid.Points[0][rows-1].Fixed = true
	for i := range grid.Lines {
		if i < rows-1+(rows*2-1)*(cols/3) {
			grid.Lines[i].Color = colornames.Green
		} else if i < rows-1+(rows-1+rows)*(cols/3*2) {
			grid.Lines[i].Color = colornames.White
		} else {
			grid.Lines[i].Color = colornames.Red
		}
	}
}

func update(click *pixel.Vec) {
	if click != nil {
		v := verlet.Vector(*click)
		grid.Head.Position = &v
	}

	grid.Update(6)
}

func draw(imd *imdraw.IMDraw) {
	for _, l := range grid.Lines {
		imd.Color = l.Color
		imd.Push(pixel.Vec(*l.A.Position), pixel.Vec(*l.B.Position))
		imd.Line(2)
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
