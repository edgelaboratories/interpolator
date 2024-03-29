package interpolator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPiecewiseConstant(t *testing.T) {
	_, err := NewPiecewiseConstant(testLinearXYs)
	require.NoError(t, err)
}

func TestNewPiecewiseConstantEmptyXYs(t *testing.T) {
	_, err := NewPiecewiseConstant(XYs{})
	require.Error(t, err)
}

func TestPiecewiseConstantValue_SinglePoint(t *testing.T) {
	const tol = 1e-15

	interpolator, err := NewPiecewiseConstant(XYs{
		{
			X: 0.0,
			Y: 0.0,
		},
	})
	require.NoError(t, err)

	assert.InDelta(t, 0.0, interpolator.Value(-1.0), tol)
	assert.InDelta(t, 0.0, interpolator.Value(0.0), tol)
	assert.InDelta(t, 0.0, interpolator.Value(1.0), tol)

	assert.InDelta(t, 0.0, interpolator.Gradient(-1.0), tol)
	assert.InDelta(t, 0.0, interpolator.Gradient(0.0), tol)
	assert.InDelta(t, 0.0, interpolator.Gradient(1.0), tol)
}

func TestPiecewiseConstantValue(t *testing.T) {
	tolerance := 1.0e-8
	interpolator, err := NewPiecewiseConstant(testLinearXYs)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			"LeftExtrapolation",
			-1.2,
			testLinearFunc(0.0),
		},
		{
			"Interpolation1",
			0.3,
			testLinearFunc(0.0),
		},
		{
			"Interpolation2",
			0.7,
			testLinearFunc(0.5),
		},
		{
			"Interpolation3",
			1.2,
			testLinearFunc(1.0),
		},
		{
			"Interpolation4",
			1.6,
			testLinearFunc(1.5),
		},
		{
			"RightExtrapolation",
			6.0,
			testLinearFunc(2.0),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.InEpsilon(t, tc.expected, interpolator.Value(tc.input), tolerance)
		})
	}
}

func ExamplePiecewiseConstant_Value() {
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
	interp, err := NewPiecewiseConstant(xys)
	if err != nil {
		return
	}
	fmt.Println(interp.Value(0.75))
	// Output: 1
}

func TestPiecewiseConstantGradient(t *testing.T) {
	tolerance := 1.0e-8
	interpolator, err := NewPiecewiseConstant(testLinearXYs)
	require.NoError(t, err)

	testCases := []struct {
		name  string
		input float64
	}{
		{
			"LeftExtrapolation",
			-1.2,
		},
		{
			"Interpolation1",
			0.3,
		},
		{
			"Interpolation2",
			0.7,
		},
		{
			"Interpolation3",
			1.2,
		},
		{
			"Interpolation4",
			1.6,
		},
		{
			"RightExtrapolation",
			6.0,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.InDelta(t, 0.0, interpolator.Gradient(tc.input), tolerance)
		})
	}
}

func ExamplePiecewiseConstant_Gradient() {
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
	interp, err := NewPiecewiseConstant(xys)
	if err != nil {
		return
	}
	fmt.Println(interp.Gradient(0.75))
	// Output: 0
}
