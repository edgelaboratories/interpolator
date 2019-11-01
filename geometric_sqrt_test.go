package interpolator

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testExpSqrtXYs = XYs{
	{
		X: 0.0,
		Y: math.Exp(0.0),
	},
	{
		X: 0.5,
		Y: math.Exp(math.Sqrt(0.5)),
	},
	{
		X: 1.0,
		Y: math.Exp(2.0 * math.Sqrt(0.5)),
	},
	{
		X: 1.5,
		Y: math.Exp(3.0 * math.Sqrt(0.5)),
	},
	{
		X: 2.0,
		Y: math.Exp(4.0 * math.Sqrt(0.5)),
	},
}

func TestNewGeometricSqrt(t *testing.T) {
	_, err := NewGeometricSqrt(testLinearXYs)
	assert.NoError(t, err)
}

func TestNewGeometricSqrtEmptyXYs(t *testing.T) {
	_, err := NewGeometricSqrt(XYs{})
	assert.Error(t, err)
}

func TestNewGeometricSqrtInsufficientXYs(t *testing.T) {
	_, err := NewGeometricSqrt(XYs{
		{
			X: 0.0,
			Y: 1.0,
		},
	})
	assert.Error(t, err)
}

func TestNewGeometricSqrtInvalidXYs(t *testing.T) {
	_, err := NewGeometricSqrt(XYs{
		{
			X: 0.0,
			Y: epsilon / 2.0,
		},
		{
			X: 1.0,
			Y: 1.0,
		},
	})
	assert.Error(t, err)
}

func TestGeometricSqrtValue(t *testing.T) {
	tolerance := 1.0e-8
	interpolator, err := NewGeometricSqrt(testExpSqrtXYs)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			"LeftExtrapolation",
			-1.0,
			1.0,
		},
		{
			"Interpolation1",
			0.3,
			math.Exp(math.Sqrt(0.3)),
		},
		{
			"Interpolation2",
			0.7,
			math.Exp(math.Sqrt(0.2) + math.Sqrt(0.5)),
		},
		{
			"Interpolation3",
			1.2,
			math.Exp(math.Sqrt(0.2) + 2.0*math.Sqrt(0.5)),
		},
		{
			"Interpolation4",
			1.6,
			math.Exp(math.Sqrt(0.1) + 3.0*math.Sqrt(0.5)),
		},
		{
			"RightExtrapolation",
			6.0,
			math.Exp(4.0 * math.Sqrt(0.5)),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.InEpsilon(t, tc.expected, interpolator.Value(tc.input), tolerance)
		})
	}
}

func ExampleGeometricSqrt_Value() {
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
	interp, err := NewGeometricSqrt(xys)
	if err != nil {
		return
	}
	fmt.Printf("%0.4f\n", interp.Value(0.75))
	// Output: 1.2686
}

func TestGeometricSqrtGradient(t *testing.T) {
	tolerance := 1.0e-8
	interpolator, err := NewGeometricSqrt(testExpSqrtXYs)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			"LeftExtrapolation",
			-1.0,
			0.0,
		},
		{
			"Interpolation1",
			0.3,
			math.Sqrt(0.5) * math.Exp(math.Sqrt(0.3)) / math.Sqrt(0.3/0.5),
		},
		{
			"Interpolation2",
			0.7,
			math.Sqrt(0.5) * (math.Exp(math.Sqrt(0.2) + math.Sqrt(0.5))) / math.Sqrt(0.2/0.5),
		},
		{
			"Interpolation3",
			1.2,
			math.Sqrt(0.5) * (math.Exp(math.Sqrt(0.2) + 2.0*math.Sqrt(0.5))) / math.Sqrt(0.2/0.5),
		},
		{
			"Interpolation4",
			1.6,
			math.Sqrt(0.5) * (math.Exp(math.Sqrt(0.1) + 3.0*math.Sqrt(0.5))) / math.Sqrt(0.1/0.5),
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

func ExampleGeometricSqrt_Gradient() {
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
	interp, err := NewGeometricSqrt(xys)
	if err != nil {
		return
	}
	fmt.Printf("%0.4f\n", interp.Gradient(0.75))
	// Output: 0.6037
}

func BenchmarkGeometricSqrtValue(b *testing.B) {
	interpolator, err := NewGeometricSqrt(testExpSqrtXYs)
	require.NoError(b, err)
	var (
		x = 0.7
		y = math.Exp(math.Sqrt(0.2) + math.Sqrt(0.5))
	)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		assert.InDelta(b, y, interpolator.Value(x), 1.0e-8)
	}
}

func BenchmarkGeometricSqrtDerivative(b *testing.B) {
	interpolator, err := NewGeometricSqrt(testExpSqrtXYs)
	require.NoError(b, err)
	var (
		x = 0.7
		y = math.Sqrt(0.5) * (math.Exp(math.Sqrt(0.2) + math.Sqrt(0.5))) / math.Sqrt(0.2/0.5)
	)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		assert.InDelta(b, y, interpolator.Gradient(x), 1.0e-8)
	}
}
