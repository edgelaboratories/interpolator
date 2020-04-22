package interpolator

import "fmt"

// PiecewiseLinear is a classic piecewise linear interpolator.
type PiecewiseLinear struct {
	xys XYs
}

// NewPiecewiseLinear builds a piecewise linear interpolator.
func NewPiecewiseLinear(xys XYs) (*PiecewiseLinear, error) {
	if l := len(xys); l < 1 {
		return nil, fmt.Errorf("at least 1 points is required to build a piecewise linear interpolator, but got %d", l)
	}
	return &PiecewiseLinear{
		xys: xys,
	}, nil
}

// Value compute the value of f(x) based on piecewise linear interpolation.
func (interp PiecewiseLinear) Value(x float64) float64 {
	if n := len(interp.xys); n == 1 {
		// In case a single data point is provided, assume a constant curve
		return interp.xys[0].Y
	}
	p1, p2 := interp.xys.Interval(x)
	lambda := (x - p1.X) / (p2.X - p1.X)
	return p1.Y*(1.0-lambda) + p2.Y*lambda
}

// Gradient computes the gradient of f(x) based on linear interpolation.
func (interp PiecewiseLinear) Gradient(x float64) float64 {
	if n := len(interp.xys); n == 1 {
		// In case a single data point is provided, assume a constant curve
		return 0.0
	}
	p1, p2 := interp.xys.Interval(x)
	return (p2.Y - p1.Y) / (p2.X - p1.X)
}
