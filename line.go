// Package verlet contains the basic units to build more complex verlet systems.
package verlet

// Line constrains two points.
type Line struct {
	A   *Point
	B   *Point
	Len float64
}

// Update tries to adjust the positiong of the points
// of a line according to its length.
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
