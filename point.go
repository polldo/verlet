package verlet

import (
	"math"
)

type PointOpt func(*Point)

func (p *Point) SetOptions(opts ...PointOpt) {
	for _, opt := range opts {
		opt(p)
	}
}

func Fix() PointOpt {
	return func(p *Point) {
		p.Fixed = true
	}
}

func Radius(r float64) PointOpt {
	return func(p *Point) {
		p.Radius = r
	}
}

type Point struct {
	Fixed       bool
	Radius      float64
	Position    Vector
	OldPosition Vector
}

func (p *Point) Distance(other *Point) float64 {
	diff := p.Position.Sub(other.Position)
	dist := math.Sqrt(diff.X*diff.X + diff.Y*diff.Y)
	return dist
}

func (p *Point) Update(friction float64, gravity Vector) {
	if p.Fixed {
		return
	}
	// Retrieve velocity as delta of space
	vel := p.Position.Sub(p.OldPosition)
	// Apply friction
	vel = vel.Scale(friction)
	p.OldPosition = p.Position
	// Apply velocity
	p.Position = p.Position.Add(vel)
	// Apply gravity
	p.Position = p.Position.Add(gravity)
}

func (p *Point) Bounds(bound Vector) {
	vel := p.Position.Sub(p.OldPosition)
	// Check bounds
	if p.Position.X > bound.X-p.Radius {
		p.Position.X = bound.X - p.Radius
		p.OldPosition.X = p.Position.X + vel.X
	} else if p.Position.X < 0+p.Radius {
		p.Position.X = 0 + p.Radius
		p.OldPosition.X = p.Position.X + vel.X
	}
	if p.Position.Y > bound.Y-p.Radius {
		p.Position.Y = bound.Y - p.Radius
		p.OldPosition.Y = p.Position.Y + vel.Y
	} else if p.Position.Y < 0+p.Radius {
		p.Position.Y = 0 + p.Radius
		p.OldPosition.Y = p.Position.Y + vel.Y
	}
}
