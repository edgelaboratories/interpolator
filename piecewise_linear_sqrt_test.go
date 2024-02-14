package interpolator

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testLinearSqrtXYs = XYs{
	{
		X: 0.0,
		Y: testLinearFunc(0.0),
	},
	{
		X: 0.5,
		Y: testLinearFunc(0.5),
	},
	{
		X: 1.0,
		Y: testLinearFunc(1.0),
	},
	{
		X: 1.5,
		Y: testLinearFunc(1.5),
	},
	{
		X: 2.0,
		Y: testLinearFunc(2.0),
	},
}

func TestNewPiecewiseLinearSqrt(t *testing.T) {
	_, err := NewPiecewiseLinearSqrt(testLinearXYs)
	require.NoError(t, err)
}

func TestNewPiecewiseLinearSqrtEmptyXYs(t *testing.T) {
	_, err := NewPiecewiseLinearSqrt(XYs{})
	require.Error(t, err)
}

func TestNewPiecewiseLinearSqrtSinglePoint(t *testing.T) {
	const tol = 1e-15

	interpolator, err := NewPiecewiseLinearSqrt(XYs{
		{
			X: 0.0,
			Y: 1.0,
		},
	})
	require.NoError(t, err)

	assert.InDelta(t, 1.0, interpolator.Value(-1.0), tol)
	assert.InDelta(t, 1.0, interpolator.Value(0.0), tol)
	assert.InDelta(t, 1.0, interpolator.Value(1.0), tol)

	assert.InDelta(t, 0.0, interpolator.Gradient(-1.0), tol)
	assert.InDelta(t, 0.0, interpolator.Gradient(0.0), tol)
	assert.InDelta(t, 0.0, interpolator.Gradient(1.0), tol)
}

func TestPiecewiseLinearSqrtValue(t *testing.T) {
	tolerance := 1.0e-8
	interpolator, err := NewPiecewiseLinearSqrt(testLinearSqrtXYs)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			"LeftExtrapolation",
			-1.2,
			1.2,
		},
		{
			"Interpolation1",
			0.3,
			1.2 * (math.Sqrt(0.15) + 1.0),
		},
		{
			"Interpolation2",
			0.7,
			1.2 * (math.Sqrt(0.10) + 1.5),
		},
		{
			"Interpolation3",
			1.2,
			1.2 * (math.Sqrt(0.10) + 2.0),
		},
		{
			"Interpolation4",
			1.6,
			1.2 * (math.Sqrt(0.05) + 2.5),
		},
		{
			"RightExtrapolation",
			6.0,
			3.6,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.InEpsilon(t, tc.expected, interpolator.Value(tc.input), tolerance)
		})
	}
}

func ExamplePiecewiseLinearSqrt_Value() {
	xys := XYs{
		{
			X: 0.0,
			Y: 1.2,
		},
		{
			X: 0.5,
			Y: 1.0,
		},
		{
			X: 1.0,
			Y: 1.4,
		},
	}
	interp, err := NewPiecewiseLinearSqrt(xys)
	if err != nil {
		return
	}
	fmt.Println(interp.Value(1.5))
	// Output: 1.4
}

func TestPiecewiseLinearSqrtValueConstantXYs(t *testing.T) {
	tolerance := 1.0e-8
	constant := 2.0
	interpolator, err := NewPiecewiseLinearSqrt(XYs{
		{
			X: 0.0,
			Y: constant,
		},
		{
			X: 1.0,
			Y: constant,
		},
	})
	require.NoError(t, err)
	assert.InEpsilon(t, constant, interpolator.Value(-1.0), tolerance)
	assert.InEpsilon(t, constant, interpolator.Value(0.5), tolerance)
	assert.InEpsilon(t, constant, interpolator.Value(2.0), tolerance)
}

func TestPiecewiseLinearSqrtGradientLinearXYs(t *testing.T) {
	tolerance := 1.0e-8
	interpolator, err := NewPiecewiseLinearSqrt(testLinearXYs)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			"LeftExtrapolation",
			-1.2,
			0.0,
		},
		{
			"Interpolation1",
			0.3,
			0.5 * 1.2 * 0.5 / math.Sqrt(0.3*0.5),
		},
		{
			"Interpolation2",
			0.7,
			0.5 * 1.2 * 0.5 / math.Sqrt(0.2*0.5),
		},
		{
			"Interpolation3",
			1.2,
			0.5 * 1.2 * 0.5 / math.Sqrt(0.2*0.5),
		},
		{
			"Interpolation4",
			1.6,
			0.5 * 1.2 * 0.5 / math.Sqrt(0.1*0.5),
		},
		{
			"RightExtrapolation",
			6.0,
			0.0,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.InDelta(t, tc.expected, interpolator.Gradient(tc.input), tolerance)
		})
	}
}

func ExamplePiecewiseLinearSqrt_Gradient() {
	xys := XYs{
		{
			X: 0.0,
			Y: 1.2,
		},
		{
			X: 0.5,
			Y: 1.0,
		},
		{
			X: 1.0,
			Y: 1.4,
		},
	}
	interp, err := NewPiecewiseLinearSqrt(xys)
	if err != nil {
		return
	}
	fmt.Println(interp.Gradient(1.5))
	// Output: 0
}
