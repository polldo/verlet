package verlet

type Vector struct {
	X, Y float64
}

func (v *Vector) Add(other *Vector) *Vector {
	return &Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v *Vector) Sub(other *Vector) *Vector {
	return &Vector{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v *Vector) Scale(scalar float64) *Vector {
	return &Vector{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}
