package interpolator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testLinearFunc(x float64) float64 {
	return 1.2 * (x + 1.0)
}

var testLinearXYs = XYs{
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

func TestXYsInterval(t *testing.T) {
	data := testLinearXYs

	testCases := []struct {
		name  string
		input float64
		x1    XY
		x2    XY
	}{
		{
			"LeftExtrapolation",
			-1.0,
			data[0],
			data[1],
		},
		{
			"Interpolation1",
			0.3,
			data[0],
			data[1],
		},
		{
			"Interpolation2",
			0.7,
			data[1],
			data[2],
		},
		{
			"Interpolation3",
			1.2,
			data[2],
			data[3],
		},
		{
			"Interpolation4",
			1.6,
			data[3],
			data[4],
		},
		{
			"RightExtrapolation",
			6.0,
			data[3],
			data[4],
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			x1, x2 := data.Interval(tc.input)

			assert.Equal(t, tc.x1, x1)
			assert.Equal(t, tc.x2, x2)
		})
	}
}

func ExampleXYs_Interval() {
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
	fmt.Println(xys.Interval(0.75))
	// Output: {0.5 1} {1 1.4}
}
