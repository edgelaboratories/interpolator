package interpolator

import "sort"

// XY represent a 2-dimensional data point.
type XY struct {
	X float64
	Y float64
}

// XYs represents a slice of data points.
type XYs []XY

// Interval returns the bracketing points around x for a given XYs.
// The `xys` must be ordered and have unique abscissas.
func (xys XYs) Interval(x float64) (XY, XY) {
	n := len(xys)
	if x <= xys[0].X {
		return xys[0], xys[1]
	}
	if x >= xys[n-1].X {
		return xys[n-2], xys[n-1]
	}
	upperBound := sort.Search(n, func(i int) bool { return xys[i].X > x })
	return xys[upperBound-1], xys[upperBound]
}
