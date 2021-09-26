package main

import (
	"math/rand"
	"time"

	"github.com/polldo/verlet"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Worm struct {
	*verlet.Rope
	velocity verlet.Vector
	moveTime int64
}

func NewWorm() *Worm {
	w := &Worm{}
	p := &verlet.VerletParams{
		Gravity:  verlet.Vector{X: 0, Y: 0},
		Bound:    verlet.Vector{X: windowWidth, Y: windowHeight},
		Friction: 0.9,
	}
	w.Rope = verlet.NewRope(10, 10, p)
	w.moveTime = time.Now().UnixNano() / int64(time.Millisecond)
	return w
}

func (w *Worm) Update() {
	w.Head.Position = w.Head.Position.Add(&w.velocity)
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
		imd.Push(pixel.Vec(*l.A.Position), pixel.Vec(*l.B.Position))
		imd.Line(5)
	}
	imd.Color = colornames.Orange
	imd.Push(pixel.V(w.Head.Position.X, w.Head.Position.Y))
	imd.Circle(8, 0)
}
