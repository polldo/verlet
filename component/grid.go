package component

import "github.com/polldo/verlet"

type Grid struct {
	*verlet.Verlet
	Origin     *verlet.Point
	Rows, Cols int
}

func NewGrid(cols, rows int, distance float64, opts ...verlet.Opt) *Grid {
	g := &Grid{
		Verlet: verlet.New(opts...),
		Rows:   rows,
		Cols:   cols,
	}

	rad := 2.
	x := g.Verlet.Bound.X / 2
	y := g.Verlet.Bound.Y / 2
	for i := 0; i < cols; i++ {

		if i == 0 {
			g.Origin = g.NewPoint(x, y, verlet.Radius(rad), verlet.Fix())
		} else {
			g.NewPoint(x+float64(i)*distance, y, verlet.Radius(rad))
			g.NewLine(g.Extract(i, 0), g.Extract(i-1, 0))
		}

		for j := 1; j < rows; j++ {
			g.NewPoint(x+float64(i)*distance, y-float64(j)*distance, verlet.Radius(rad))
			g.NewLine(g.Extract(i, j), g.Extract(i, j-1))
			if i > 0 {
				g.NewLine(g.Extract(i, j), g.Extract(i-1, j))
			}
		}
	}

	return g
}

func (g *Grid) Extract(col, row int) *verlet.Point {
	idx := g.MatrixToArray(col, row)
	return g.Points[idx]
}

func (g *Grid) MatrixToArray(col, row int) int {
	return col*g.Rows + row
}
