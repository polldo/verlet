package verlet

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

type Verlet struct {
	Points   []*Point
	Lines    []*Line
	Gravity  Vector
	Bound    Vector
	Friction float64
}

func New(opts ...Opt) *Verlet {
	v := &Verlet{
		Gravity:  Vector{X: 0.1, Y: -0.1},
		Bound:    Vector{X: 100, Y: 100},
		Friction: 0,
	}
	v.SetOptions(opts...)
	return v
}

func (v *Verlet) SetOptions(opts ...Opt) {
	for _, opt := range opts {
		opt(v)
	}
}

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

func (v *Verlet) NewLine(a, b *Point) *Line {
	ln := &Line{
		A:   a,
		B:   b,
		Len: a.Distance(b),
	}
	v.Lines = append(v.Lines, ln)
	return ln
}

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
