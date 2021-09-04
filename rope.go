package verlet

import (
	"golang.org/x/image/colornames"
)

type Rope struct {
	Points []*Point
	Lines  []*Line
	Head   *Point
}

func NewRope(units int, distance float64, params *Params) *Rope {
	r := &Rope{}

	rad := 8.
	r.Head = NewPoint(params.BoundX/2, params.BoundY/2, rad, true, colornames.Orange, params)
	r.Points = append(r.Points, r.Head)

	for i := 1; i < units; i++ {
		r.Points = append(r.Points, NewPoint(params.BoundX/2, params.BoundY/2-float64(i)*distance, rad, false, colornames.Orange, params))
		r.Lines = append(r.Lines, NewLine(r.Points[i], r.Points[i-1], colornames.Orange))
	}
	return r
}

func (r *Rope) Update(count int) {
	for _, p := range r.Points {
		p.Update()
	}
	for i := 0; i < count; i++ {
		for _, l := range r.Lines {
			l.Update()
		}
	}
}
