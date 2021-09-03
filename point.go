package verlet

import (
	"image/color"
	"math"
)

type Point struct {
	Color       color.RGBA
	Fixed       bool
	Radius      float64
	Position    *Vector
	OldPosition *Vector
	Bound       *Vector
	Gravity     *Vector
	Friction    float64
}

type Params struct {
	GravityX float64
	GravityY float64
	BoundX   float64
	BoundY   float64
	Friction float64
}

func NewPoint(x, y float64, radius float64, fixed bool, c color.RGBA, params *Params) *Point {
	return &Point{
		Color:       c,
		Fixed:       fixed,
		Radius:      radius,
		Position:    &Vector{X: x, Y: y},
		OldPosition: &Vector{X: x, Y: y},
		Bound:       &Vector{X: params.BoundX, Y: params.BoundY},
		Gravity:     &Vector{X: params.GravityX, Y: params.GravityY},
		Friction:    params.Friction,
	}
}

func (p *Point) Distance(other *Point) float64 {
	diff := p.Position.Sub(other.Position)
	dist := math.Sqrt(diff.X*diff.X + diff.Y*diff.Y)
	return dist
}

func (p *Point) Update() {
	if p.Fixed {
		return
	}
	// Retrieve velocity as delta of space
	vel := p.Position.Sub(p.OldPosition)
	// Apply friction
	vel = vel.Scale(p.Friction)
	p.OldPosition = p.Position
	// Apply velocity
	p.Position = p.Position.Add(vel)
	// Apply gravity
	p.Position = p.Position.Add(p.Gravity)

	// Check bounds
	if p.Position.X > p.Bound.X-p.Radius {
		p.Position.X = p.Bound.X - p.Radius
		p.OldPosition.X = p.Position.X + vel.X
	} else if p.Position.X < 0+p.Radius {
		p.Position.X = 0 + p.Radius
		p.OldPosition.X = p.Position.X + vel.X
	}
	if p.Position.Y > p.Bound.Y-p.Radius {
		p.Position.Y = p.Bound.Y - p.Radius
		p.OldPosition.Y = p.Position.Y + vel.Y
	} else if p.Position.Y < 0+p.Radius {
		p.Position.Y = 0 + p.Radius
		p.OldPosition.Y = p.Position.Y + vel.Y
	}
}
