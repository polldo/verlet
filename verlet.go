package verlet

type VerletParams struct {
	Gravity  Vector
	Bound    Vector
	Friction float64
}

type Verlet struct {
	VerletParams
	Points []*Point
	Lines  []*Line
}

func New(params *VerletParams) *Verlet {
	return &Verlet{
		VerletParams: *params,
	}
}

func (v *Verlet) NewPoint(x, y float64, radius float64, fixed bool) *Point {
	p := &Point{
		Fixed:       fixed,
		Radius:      radius,
		Position:    Vector{X: x, Y: y},
		OldPosition: Vector{X: x, Y: y},
	}

	v.Points = append(v.Points, p)
	return p
}

func (v *Verlet) NewLine(a, b *Point) *Line {
	l := a.Distance(b)
	ln := &Line{
		A:   a,
		B:   b,
		Len: l,
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
