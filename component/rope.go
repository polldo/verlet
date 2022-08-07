// Package component leverages the basic verlet units to
// implement more complex components.
// It can be seen as an example to make custom components.
package component

import "github.com/polldo/verlet"

// Rope is a component that represents a simple rope.
// It's composed of several concatenated lines.
// A fixed point is used as the rope's head.
type Rope struct {
	*verlet.Verlet
	Head *verlet.Point
}

// NewRope builds and returns a new rope.
func NewRope(units int, distance float64, opts ...verlet.Opt) *Rope {
	r := &Rope{
		Verlet: verlet.New(opts...),
	}

	// Center the rope.
	x := r.Verlet.Bound.X / 2
	y := r.Verlet.Bound.Y / 2
	r.Head = r.NewPoint(x, y, verlet.Fix(), verlet.Radius(8.0))

	// Build and concatenate lines.
	for i := 1; i < units; i++ {
		r.NewPoint(x, y-float64(i)*distance)
		r.NewLine(r.Points[i], r.Points[i-1])
	}
	return r
}
