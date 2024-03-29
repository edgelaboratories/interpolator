package interpolator

import (
	"errors"
	"fmt"
	"math"
)

// GeometricSqrt performs a geometric interpolation with respect to the square root of the abscissae.
type GeometricSqrt struct {
	xys XYs
}

// NewGeometricSqrt builds a geometric sqrt interpolator.
// The input `xys` must be ordered, have unique abscissas
// and positive ordinates.
func NewGeometricSqrt(xys XYs) (*GeometricSqrt, error) {
	if l := len(xys); l < 1 {
		return nil, fmt.Errorf("at least 1 points is required to build a geometric sqrt interpolator, but got %d", l)
	}

	for _, xy := range xys {
		if xy.Y < epsilon {
			return nil, errors.New("input xys must have non-negative ordinates")
		}
	}

	return &GeometricSqrt{
		xys: xys,
	}, nil
}

// Value compute the value of f(x) based on geometric sqrt interpolation with flat extrapolation.
func (interp GeometricSqrt) Value(x float64) float64 {
	if n := len(interp.xys); n == 1 {
		return interp.xys[0].Y
	}

	p1, p2 := interp.xys.Interval(x)
	if x <= p1.X {
		return p1.Y
	}
	if x >= p2.X {
		return p2.Y
	}

	lambda := math.Sqrt((x - p1.X) / (p2.X - p1.X))

	return math.Pow(p1.Y, (1.0-lambda)) * math.Pow(p2.Y, lambda)
}

// Gradient computes the gradient of f(x) based on geometric sqrt interpolation.
func (interp GeometricSqrt) Gradient(x float64) float64 {
	if n := len(interp.xys); n == 1 {
		return 0.0
	}

	p1, p2 := interp.xys.Interval(x)
	if x <= p1.X || x >= p2.X {
		return 0.0
	}

	h := p2.X - p1.X
	lambda := math.Sqrt((x - p1.X) / h)

	return 0.5 * math.Log(p2.Y/p1.Y) * math.Pow(p1.Y, (1.0-lambda)) * math.Pow(p2.Y, lambda) / (lambda * h)
}
