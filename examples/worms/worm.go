package main

import (
	"math/rand"
	"time"

	"github.com/polldo/verlet"
	"github.com/polldo/verlet/component"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Worm struct {
	*component.Rope
	velocity verlet.Vector
	moveTime int64
}

func NewWorm() *Worm {
	w := &Worm{}
	w.Rope = component.NewRope(
		10,
		10,
		verlet.Gravity(0, 0),
		verlet.Bound(windowWidth, windowHeight),
		verlet.Friction(0.9),
	)
	w.moveTime = time.Now().UnixNano() / int64(time.Millisecond)
	return w
}

func (w *Worm) Update() {
	w.Head.Position = w.Head.Position.Add(w.velocity)
	w.Rope.Update(30)
	if time.Now().UnixNano()/int64(time.Millisecond)-w.moveTime > 1000 {
		vx, vy := rand.Float64()*10-5, rand.Float64()*10-5
		w.velocity = verlet.Vector{X: vx, Y: vy}
		w.moveTime = time.Now().UnixNano() / int64(time.Millisecond)
	}
}

func (w *Worm) Draw(imd *imdraw.IMDraw) {
	for _, l := range w.Lines {
		imd.Color = colornames.Rosybrown
		imd.Push(pixel.Vec(l.A.Position), pixel.Vec(l.B.Position))
		imd.Line(5)
	}
	imd.Color = colornames.Orange
	imd.Push(pixel.V(w.Head.Position.X, w.Head.Position.Y))
	imd.Circle(8, 0)
}
