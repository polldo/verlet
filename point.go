// Package verlet contains the basic units to build more complex verlet systems.
package verlet

import (
	"math"
)

// PointOpt is the type representing a functional option
// for verlet points.
type PointOpt func(*Point)

// SetOptions configures the point p with the passed
// functional options.
func (p *Point) SetOptions(opts ...PointOpt) {
	for _, opt := range opts {
		opt(p)
	}
}

// Fix returns a functional option to make a point fixed.
// Fixed points cannot move.
func Fix() PointOpt {
	return func(p *Point) {
		p.Fixed = true
	}
}

// Radius returns a functional option that sets the radius
// of points.
func Radius(r float64) PointOpt {
	return func(p *Point) {
		p.Radius = r
	}
}

// Point is the atom of every verlet system.
type Point struct {
	Fixed       bool // Fixed points cannot move.
	Radius      float64
	Position    Vector
	OldPosition Vector
}

// Distance returns the distance between two points.
func (p *Point) Distance(other *Point) float64 {
	diff := p.Position.Sub(other.Position)
	dist := math.Sqrt(diff.X*diff.X + diff.Y*diff.Y)
	return dist
}

// Update applies the given forces to the point and
// updates its point position accordingly.
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

// Bounds handles collisions of the point with the boundaries
// of the system and regulates the point velocity accordingly.
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
