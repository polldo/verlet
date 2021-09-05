package verlet

type Rope struct {
	*Verlet
	Head *Point
}

func NewRope(units int, distance float64, params *VerletParams) *Rope {
	r := &Rope{
		Verlet: New(params),
	}

	rad := 8.
	x := params.Bound.X / 2
	y := params.Bound.Y / 2
	r.Head = r.NewPoint(x, y, rad, true)

	for i := 1; i < units; i++ {
		r.NewPoint(x, y-float64(i)*distance, rad, false)
		r.NewLine(r.Points[i], r.Points[i-1])
	}

	return r
}
