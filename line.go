package main

import (
	"image/color"

	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Line struct {
	A     *Point
	B     *Point
	Len   float64
	Color color.RGBA
}

func NewLine(a, b *Point) *Line {
	l := a.Distance(b)
	return &Line{
		A:     a,
		B:     b,
		Len:   l,
		Color: colornames.Black,
	}
}

func (l *Line) Update() {
	distance := l.A.Distance(l.B)
	dL := (l.Len - distance) / distance / 2
	offset := l.A.Position.Sub(l.B.Position).Scaled(dL)
	if !l.A.Fixed {
		l.A.Position = l.A.Position.Add(offset)
	}
	if !l.B.Fixed {
		l.B.Position = l.B.Position.Sub(offset)
	}
}

func (l *Line) Draw(imd *imdraw.IMDraw) {
	imd.Color = l.Color
	imd.Push(l.A.Position, l.B.Position)
	imd.Line(2)
	// fmt.Println(l.A.Position, l.B.Position)
}
