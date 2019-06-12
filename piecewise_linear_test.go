package interpolator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPiecewiseLinear(t *testing.T) {
	_, err := NewPiecewiseLinear(testLinearXYs)
	assert.NoError(t, err)
}

func TestNewPiecewiseLinearEmptyXYs(t *testing.T) {
	_, err := NewPiecewiseLinear(XYs{})
	assert.Error(t, err)
}

func TestNewPiecewiseLinearInsufficientXYs(t *testing.T) {
	_, err := NewPiecewiseLinear(XYs{
		{
			X: 0.0,
			Y: 1.0,
		},
	})
	assert.Error(t, err)
}

func TestPiecewiseLinearValue(t *testing.T) {
	tolerance := 1.0e-8
	interpolator, err := NewPiecewiseLinear(testLinearXYs)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			"LeftExtrapolation",
			-1.2,
			testLinearFunc(-1.2),
		},
		{
			"Interpolation1",
			0.3,
			testLinearFunc(0.3),
		},
		{
			"Interpolation2",
			0.7,
			testLinearFunc(0.7),
		},
		{
			"Interpolation3",
			1.2,
			testLinearFunc(1.2),
		},
		{
			"Interpolation4",
			1.6,
			testLinearFunc(1.6),
		},
		{
			"RightExtrapolation",
			6.0,
			testLinearFunc(6.0),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.InEpsilon(t, tc.expected, interpolator.Value(tc.input), tolerance)
		})
	}
}

func ExamplePiecewiseLinear_Value() {
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
	interp, err := NewPiecewiseLinear(xys)
	if err != nil {
		return
	}
	fmt.Println(interp.Value(0.75))
	// Output: 1.2
}

func TestPiecewiseLinearValueConstantXYs(t *testing.T) {
	tolerance := 1.0e-8
	constant := 2.0
	interpolator, err := NewPiecewiseLinear(XYs{
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

func TestPiecewiseLinearGradientLinearXYs(t *testing.T) {
	tolerance := 1.0e-8
	interpolator, err := NewPiecewiseLinear(testLinearXYs)
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
			1.2,
		},
		{
			"Interpolation2",
			0.7,
			1.2,
		},
		{
			"Interpolation3",
			1.2,
			1.2,
		},
		{
			"Interpolation4",
			1.6,
			1.2,
		},
		{
			"RightExtrapolation",
			6.0,
			1.2,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.InEpsilon(t, tc.expected, interpolator.Gradient(tc.input), tolerance)
		})
	}
}

func ExamplePiecewiseLinear_Gradient() {
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
	interp, err := NewPiecewiseLinear(xys)
	if err != nil {
		return
	}
	fmt.Printf("%0.1f\n", interp.Gradient(0.75))
	// Output: 0.8
}
