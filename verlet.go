// Package verlet contains the basic units to build more complex verlet systems.
package verlet

// Opt represents a functional option for verlet.
type Opt func(*Verlet)

// Gravity sets the gravity option for verlet.
func Gravity(x, y float64) Opt {
	return func(v *Verlet) {
		v.Gravity = Vector{X: x, Y: y}
	}
}

// Bound sets the bounds option for verlet.
func Bound(x, y float64) Opt {
	return func(v *Verlet) {
		v.Bound = Vector{X: x, Y: y}
	}
}

// Friction sets the friction option for verlet.
func Friction(f float64) Opt {
	return func(v *Verlet) {
		v.Friction = f
	}
}

// Verlet is the main unit of this library.
// It can be used to build complex shapes starting
// from lines and points.
// Its physics can be manipulated by properly
// setting its properties.
//
// It exposes all its points and line so that
// it can be iterated for rendering.
// TODO: Add an 'Apply' function that applies a
// certain function to all points of this system.
type Verlet struct {
	Points   []*Point
	Lines    []*Line
	Gravity  Vector  // Gravity force that acts on all points of this unit.
	Bound    Vector  // Bound is the size of this system (boundaries).
	Friction float64 // Resistance of verlet points.
}

// New builds a new verlet unit and returns it.
func New(opts ...Opt) *Verlet {
	v := &Verlet{
		Gravity:  Vector{X: 0.1, Y: -0.1},
		Bound:    Vector{X: 100, Y: 100},
		Friction: 0,
	}
	v.SetOptions(opts...)
	return v
}

// SetOptions reconfigures the system with the passed
// functional options.
func (v *Verlet) SetOptions(opts ...Opt) {
	for _, opt := range opts {
		opt(v)
	}
}

// NewPoint add a new point to this verlet system and returns its pointer
// so that it can be manipulated, stored or used to build new lines.
func (v *Verlet) NewPoint(x, y float64, opts ...PointOpt) *Point {
	p := &Point{
		Fixed:       false,
		Radius:      5.0,
		Position:    Vector{X: x, Y: y},
		OldPosition: Vector{X: x, Y: y},
	}
	p.SetOptions(opts...)
	v.Points = append(v.Points, p)
	return p
}

// NewLine add a new line to this verlet system and returns it.
func (v *Verlet) NewLine(a, b *Point) *Line {
	ln := &Line{
		A:   a,
		B:   b,
		Len: a.Distance(b),
	}
	v.Lines = append(v.Lines, ln)
	return ln
}

// Update is the function that animates the entire system.
// It applies all the known forces to the points and updates
// their position accordingly.
// The count value could be tweaked to change the
// precision of the system simulation.
func (v *Verlet) Update(count int) {
	for _, p := range v.Points {
		p.Update(v.Friction, v.Gravity)
		p.Bounds(v.Bound)
	}
	for i := 0; i < count; i++ {
		for _, l := range v.Lines {
			l.Update()
		}
	}
}
