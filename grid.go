package verlet

import (
	"golang.org/x/image/colornames"
)

type Grid struct {
	Points [][]*Point
	Lines  []*Line
	Head   *Point
	Tail   *Point
}

func NewGrid(cols, rows int, distance float64, params *Params) *Grid {
	g := &Grid{}

	g.Points = make([][]*Point, cols)

	rad := 2.
	for i := 0; i < cols; i++ {

		if i == 0 {
			g.Head = NewPoint(params.BoundX/2+float64(i)*distance, params.BoundY/2, rad, true, colornames.Orange, params)
			g.Points[i] = append(g.Points[i], g.Head)
		} else {
			g.Points[i] = append(g.Points[i], NewPoint(params.BoundX/2+float64(i)*distance, params.BoundY/2, rad, false, colornames.Orange, params))
			g.Lines = append(g.Lines, NewLine(g.Points[i][0], g.Points[i-1][0], colornames.Blue))
		}

		for j := 1; j < rows; j++ {
			g.Points[i] = append(g.Points[i], NewPoint(params.BoundX/2+float64(i)*distance, params.BoundY/2-float64(j)*distance, rad, false, colornames.Orange, params))
			g.Lines = append(g.Lines, NewLine(g.Points[i][j], g.Points[i][j-1], g.Points[i][j-1].Color))
			if i > 0 {
				g.Lines = append(g.Lines, NewLine(g.Points[i][j], g.Points[i-1][j], g.Points[i][j-1].Color))
			}
		}
	}

	return g
}

func (g *Grid) Update(count int) {
	for _, points := range g.Points {
		for _, p := range points {
			p.Update()
		}
	}
	for i := 0; i < count; i++ {
		for _, l := range g.Lines {
			l.Update()
		}
	}
}
