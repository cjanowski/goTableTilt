package utils

import "math"

// Vector represents a 2D vector
type Vector struct {
	X, Y float64
}

// Add adds two vectors and returns the result
func (v *Vector) Add(v2 *Vector) *Vector {
	return &Vector{X: v.X + v2.X, Y: v.Y + v2.Y}
}

// Length returns the length of the vector
func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
