package verlet

type Grid struct {
	*Verlet
	Head       *Point
	Rows, Cols int
}

func NewGrid(cols, rows int, distance float64, params *VerletParams) *Grid {
	g := &Grid{
		Verlet: New(params),
		Rows:   rows,
		Cols:   cols,
	}

	rad := 2.
	x := params.Bound.X / 2
	y := params.Bound.Y / 2
	for i := 0; i < cols; i++ {

		if i == 0 {
			g.Head = g.NewPoint(x, y, rad, true)
		} else {
			g.NewPoint(x+float64(i)*distance, y, rad, false)
			g.NewLine(g.Extract(i, 0), g.Extract(i-1, 0))
		}

		for j := 1; j < rows; j++ {
			g.NewPoint(x+float64(i)*distance, y-float64(j)*distance, rad, false)
			g.NewLine(g.Extract(i, j), g.Extract(i, j-1))
			if i > 0 {
				g.NewLine(g.Extract(i, j), g.Extract(i-1, j))
			}
		}
	}

	return g
}

func (g *Grid) Extract(col, row int) *Point {
	idx := g.MatrixToArray(col, row)
	return g.Points[idx]
}

func (g *Grid) MatrixToArray(col, row int) int {
	return col*g.Rows + row
}
