package verlet

import (
	"image/color"
)

type Line struct {
	A     *Point
	B     *Point
	Len   float64
	Color color.RGBA
}

func NewLine(a, b *Point, c color.RGBA) *Line {
	l := a.Distance(b)
	return &Line{
		A:     a,
		B:     b,
		Len:   l,
		Color: c,
	}
}

func (l *Line) Update() {
	distance := l.A.Distance(l.B)
	dL := (l.Len - distance) / distance / 2
	offset := l.A.Position.Sub(l.B.Position).Scale(dL)
	if !l.A.Fixed {
		l.A.Position = l.A.Position.Add(offset)
	}
	if !l.B.Fixed {
		l.B.Position = l.B.Position.Sub(offset)
	}
}
