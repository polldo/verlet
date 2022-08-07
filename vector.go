package verlet

// Vector is a simple struct representing a 2d math vector.
type Vector struct {
	X, Y float64
}

// Add sum 'other' to 'v' and return the result in a new vector.
func (v Vector) Add(other Vector) Vector {
	return Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

// Sub subtracts 'other' from 'v' and return the result in a new vector.
func (v Vector) Sub(other Vector) Vector {
	return Vector{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

// Scale multiplies 'v' by a scalar and returns the result in a new vector.
func (v Vector) Scale(scalar float64) Vector {
	return Vector{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}
