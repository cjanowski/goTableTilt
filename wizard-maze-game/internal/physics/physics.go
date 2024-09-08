package physics

import "math"

// Example function that uses math
func SomeFunction() float64 {
    return math.Sqrt(16) // Use math package
}

// ApplyPhysics applies basic physics to the mouse movement
func ApplyPhysics(velocity, tilt, gravity float64) float64 {
	// Example: Simple acceleration based on tilt and gravity
	return velocity + tilt*gravity
}
