package interpolator

import (
	"fmt"
	"math"
)

// epsilon is the numerical threshold under which values
// are considered too close to 0 to be interpolated.
const epsilon = 1.0e-16

// Geometric is a classic geometric interpolator.
type Geometric struct {
	xys XYs
}

// NewGeometric builds a geometric interpolator.
// The input `xys` must be ordered, have unique abscissas
// and positive ordinates.
func NewGeometric(xys XYs) (*Geometric, error) {
	if l := len(xys); l < 1 {
		return nil, fmt.Errorf("at least 1 points is required to build a geometric interpolator, but got %d", l)
	}

	for _, xy := range xys {
		if xy.Y < epsilon {
			return nil, fmt.Errorf("input xys must have non-negative ordinates")
		}
	}

	return &Geometric{
		xys: xys,
	}, nil
}

// Value compute the value of f(x) based on geometric interpolation.
func (interp Geometric) Value(x float64) float64 {
	if n := len(interp.xys); n == 1 {
		return interp.xys[0].Y
	}

	p1, p2 := interp.xys.Interval(x)
	lambda := (x - p1.X) / (p2.X - p1.X)

	return math.Pow(p1.Y, (1.0-lambda)) * math.Pow(p2.Y, lambda)
}

// Gradient computes the gradient of f(x) based on geometric interpolation.
func (interp Geometric) Gradient(x float64) float64 {
	if n := len(interp.xys); n == 1 {
		return 0.0
	}

	p1, p2 := interp.xys.Interval(x)

	return math.Log(p2.Y/p1.Y) * interp.Value(x) / (p2.X - p1.X)
}
