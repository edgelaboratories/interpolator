package interpolator

import (
	"fmt"
	"math"
)

// PiecewiseLinearSqrt performs a piecewise linear interpolation with respect to the square root of the abscissae.
type PiecewiseLinearSqrt struct {
	xys XYs
}

// NewPiecewiseLinearSqrt builds a piecewise linear sqrt interpolator with flat extrapolation.
// The input `xys` must be ordered and have unique abscissas.
func NewPiecewiseLinearSqrt(xys XYs) (*PiecewiseLinearSqrt, error) {
	if l := len(xys); l < 1 {
		return nil, fmt.Errorf("at least 1 points is required to build a piecewise linear sqrt interpolator, but got %d", l)
	}
	return &PiecewiseLinearSqrt{
		xys: xys,
	}, nil
}

// Value compute the value of f(x) based on piecewise linear interpolation with flat extrapolation.
func (interp PiecewiseLinearSqrt) Value(x float64) float64 {
	if n := len(interp.xys); n == 1 {
		// In case a single data point is provided, assume a constant curve
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
	return p1.Y*(1.0-lambda) + p2.Y*lambda
}

// Gradient computes the gradient of f(x) based on piecewise linear interpolation with flat extrapolation.
func (interp PiecewiseLinearSqrt) Gradient(x float64) float64 {
	if n := len(interp.xys); n == 1 {
		// In case a single data point is provided, assume a constant curve
		return 0.0
	}
	p1, p2 := interp.xys.Interval(x)
	if x <= p1.X || x >= p2.X {
		return 0.0
	}
	return 0.5 * (p2.Y - p1.Y) / math.Sqrt((p2.X-p1.X)*(x-p1.X))
}
