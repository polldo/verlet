package main

import (
	"image/color"

	"github.com/polldo/verlet"
	"github.com/polldo/verlet/component"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	windowWidth  = 800.
	windowHeight = 300.
)

type moveType int

const (
	movePoint moveType = iota
	movePole
)

// Flag is built on top of a grid component.
type Flag struct {
	grid *component.Grid
	mv   moveType
}

func newFlag() *Flag {
	cols, rows := 3*6, 12
	dist := 12.
	grid := component.NewGrid(cols, rows, dist, verlet.Gravity(0.1, -0.1), verlet.Bound(windowWidth, windowHeight), verlet.Friction(1))

	// Fix the points attached to the imaginary flagpole.
	grid.Extract(0, 0).Fixed = true
	grid.Extract(0, grid.Rows-1).Fixed = true
	return &Flag{grid: grid, mv: movePole}
}

func (f *Flag) update() {
	f.grid.Update(6)
}

func (f *Flag) move(pos *pixel.Vec) {
	if pos == nil {
		return
	}
	v := verlet.Vector(*pos)
	switch f.mv {

	case movePoint:
		f.grid.Origin.Position = v

	case movePole:
		up := f.grid.Extract(0, 0)
		down := f.grid.Extract(0, f.grid.Rows-1)
		dist := up.Position.Sub(down.Position)
		up.Position = v
		down.Position = up.Position.Sub(dist)
	}
}

func (f *Flag) color(idx int) color.RGBA {
	rows, cols := f.grid.Rows, f.grid.Cols
	if idx < rows-1+(rows*2-1)*(cols/3) {
		return colornames.Green
	} else if idx < rows-1+(rows*2-1)*(cols/3*2) {
		return colornames.White
	} else {
		return colornames.Red
	}
}

func draw(imd *imdraw.IMDraw, flag *Flag) {
	for i, l := range flag.grid.Lines {
		imd.Color = flag.color(i)
		imd.Push(pixel.Vec(l.A.Position), pixel.Vec(l.B.Position))
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

	flag := newFlag()

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		imd.Clear()

		if win.Pressed(pixelgl.MouseButtonLeft) {
			c := win.MousePosition()
			flag.move(&c)
		}

		flag.update()
		draw(imd, flag)

		imd.Draw(win)
		win.Update()
	}
}
