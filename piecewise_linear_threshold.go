package interpolator

import "fmt"

// PiecewiseLinearThreshold is a piecewise linear interpolator that extrapolates a threshold value.
type PiecewiseLinearThreshold struct {
	xys XYs
}

// NewPiecewiseLinearThreshold builds a piecewise linear interpolator with flat extrapolation.
// The input `xys` must be ordered and have unique abscissas.
func NewPiecewiseLinearThreshold(xys XYs) (*PiecewiseLinearThreshold, error) {
	if l := len(xys); l < 2 {
		return nil, fmt.Errorf("at least 2 points are required to build a piecewise linear threshold interpolator, but got %d", l)
	}
	return &PiecewiseLinearThreshold{
		xys: xys,
	}, nil
}

// Value compute the value of f(x) based on piecewise linear interpolation with flat extrapolation.
func (interp PiecewiseLinearThreshold) Value(x float64) float64 {
	p1, p2 := interp.xys.Interval(x)
	if x <= p1.X {
		return p1.Y
	}
	if x >= p2.X {
		return p2.Y
	}
	lambda := (x - p1.X) / (p2.X - p1.X)
	return p1.Y*(1.0-lambda) + p2.Y*lambda
}

// Gradient computes the gradient of f(x) based on piecewise linear interpolation with flat extrapolation.
func (interp PiecewiseLinearThreshold) Gradient(x float64) float64 {
	p1, p2 := interp.xys.Interval(x)
	if x <= p1.X || x >= p2.X {
		return 0.0
	}
	return (p2.Y - p1.Y) / (p2.X - p1.X)
}
