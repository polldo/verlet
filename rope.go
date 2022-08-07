package verlet

type Rope struct {
	*Verlet
	Head *Point
}

func NewRope(units int, distance float64, opts ...Opt) *Rope {
	r := &Rope{
		Verlet: New(opts...),
	}

	x := r.Verlet.Bound.X / 2
	y := r.Verlet.Bound.Y / 2
	r.Head = r.NewPoint(x, y, Fix(), Radius(8.0))

	for i := 1; i < units; i++ {
		r.NewPoint(x, y-float64(i)*distance)
		r.NewLine(r.Points[i], r.Points[i-1])
	}

	return r
}
