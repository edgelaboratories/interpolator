package interpolator

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testExpXYs = XYs{
	{
		X: 0.0,
		Y: math.Exp(0.0),
	},
	{
		X: 0.5,
		Y: math.Exp(0.5),
	},
	{
		X: 1.0,
		Y: math.Exp(1.0),
	},
	{
		X: 1.5,
		Y: math.Exp(1.5),
	},
	{
		X: 2.0,
		Y: math.Exp(2.0),
	},
}

func TestNewGeometric(t *testing.T) {
	_, err := NewGeometric(testLinearXYs)
	require.NoError(t, err)
}

func TestNewGeometricEmptyXYs(t *testing.T) {
	_, err := NewGeometric(XYs{})
	require.Error(t, err)
}

func TestNewGeometricSinglePoint(t *testing.T) {
	const tol = 1e-15

	interpolator, err := NewGeometric(XYs{
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

func TestNewGeometricInvalidXYs(t *testing.T) {
	_, err := NewGeometric(XYs{
		{
			X: 0.0,
			Y: epsilon / 2.0,
		},
		{
			X: 1.0,
			Y: 1.0,
		},
	})
	require.Error(t, err)
}

func TestGeometricValue(t *testing.T) {
	tolerance := 1.0e-8
	interpolator, err := NewGeometric(testExpXYs)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			"LeftExtrapolation",
			-1.0,
			math.Exp(-1.0),
		},
		{
			"Interpolation1",
			0.3,
			math.Exp(0.3),
		},
		{
			"Interpolation2",
			0.7,
			math.Exp(0.7),
		},
		{
			"Interpolation3",
			1.2,
			math.Exp(1.2),
		},
		{
			"Interpolation4",
			1.6,
			math.Exp(1.6),
		},
		{
			"RightExtrapolation",
			6.0,
			math.Exp(6.0),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.InEpsilon(t, tc.expected, interpolator.Value(tc.input), tolerance)
		})
	}
}

func ExampleGeometric_Value() {
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
	interp, err := NewGeometric(xys)
	if err != nil {
		return
	}
	fmt.Printf("%0.4f\n", interp.Value(0.75))
	// Output: 1.1832
}

func TestGeometricGradient(t *testing.T) {
	tolerance := 1.0e-8
	interpolator, err := NewGeometric(testExpXYs)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			"LeftExtrapolation",
			-1.0,
			math.Exp(-1.0),
		},
		{
			"Interpolation1",
			0.3,
			math.Exp(0.3),
		},
		{
			"Interpolation2",
			0.7,
			math.Exp(0.7),
		},
		{
			"Interpolation3",
			1.2,
			math.Exp(1.2),
		},
		{
			"Interpolation4",
			1.6,
			math.Exp(1.6),
		},
		{
			"RightExtrapolation",
			6.0,
			math.Exp(6.0),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.InEpsilon(t, tc.expected, interpolator.Gradient(tc.input), tolerance)
		})
	}
}

func ExampleGeometric_Gradient() {
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
	interp, err := NewGeometric(xys)
	if err != nil {
		return
	}
	fmt.Printf("%0.4f\n", interp.Gradient(0.75))
	// Output: 0.7962
}

func BenchmarkGeometricValue(b *testing.B) {
	interpolator, err := NewGeometric(testExpXYs)
	require.NoError(b, err)
	var (
		x = 0.7
		y = math.Exp(x)
	)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		assert.InDelta(b, y, interpolator.Value(x), 1.0e-8)
	}
}

func BenchmarkGeometricDerivative(b *testing.B) {
	interpolator, err := NewGeometric(testExpXYs)
	require.NoError(b, err)
	var (
		x = 0.7
		y = math.Exp(x)
	)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		assert.InDelta(b, y, interpolator.Gradient(x), 1.0e-8)
	}
}
