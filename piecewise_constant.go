package interpolator

import "errors"

// PiecewiseConstant is a classic piecewise constant cadlag interpolator.
type PiecewiseConstant struct {
	xys XYs
}

// NewPiecewiseConstant builds a piecewise constant interpolator.
// The input `xys` must be ordered and have unique abscissas.
func NewPiecewiseConstant(xys XYs) (*PiecewiseConstant, error) {
	if len(xys) < 1 {
		return nil, errors.New("at least 1 point is required to build a piecewise constant interpolator, but got 0")
	}

	return &PiecewiseConstant{
		xys: xys,
	}, nil
}

// Value compute the value of f(x) based on piecewise constant interpolation.
func (interp PiecewiseConstant) Value(x float64) float64 {
	if n := len(interp.xys); n == 1 {
		// In case a single data point is provided, assume a constant curve
		return interp.xys[n-1].Y
	}

	p1, p2 := interp.xys.Interval(x)
	if x < p2.X {
		return p1.Y
	}

	return p2.Y
}

// Gradient computes the gradient of f(x) based on piecewise constant interpolation.
func (interp PiecewiseConstant) Gradient(x float64) float64 {
	return 0.0
}
