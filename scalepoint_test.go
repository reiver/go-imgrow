package imgrow

import (
	"testing"

	"image"
)

func TestScalePoint(t *testing.T) {

	tests := []struct{
		Scalar int
		Point image.Point
		Expected image.Point
	}{
		{
			Scalar: -2,
			Point: image.Point{
				X: 5,
				Y: 7,
			},
			Expected: image.Point{
				X: -10,
				Y: -14,
			},
		},
		{
			Scalar: -1,
			Point: image.Point{
				X: 5,
				Y: 7,
			},
			Expected: image.Point{
				X: -5,
				Y: -7,
			},
		},
		{
			Scalar: 0,
			Point: image.Point{
				X: 5,
				Y: 7,
			},
			Expected: image.Point{
				X: 0,
				Y: 0,
			},
		},
		{
			Scalar: 1,
			Point: image.Point{
				X: 5,
				Y: 7,
			},
			Expected: image.Point{
				X: 5,
				Y: 7,
			},
		},
		{
			Scalar: 2,
			Point: image.Point{
				X: 5,
				Y: 7,
			},
			Expected: image.Point{
				X: 10,
				Y: 14,
			},
		},
		{
			Scalar: 3,
			Point: image.Point{
				X: 5,
				Y: 7,
			},
			Expected: image.Point{
				X: 15,
				Y: 21,
			},
		},




		{
			Scalar: 100,
			Point: image.Point{
				X: -11,
				Y: 13,
			},
			Expected: image.Point{
				X: -1100,
				Y: 1300,
			},
		},
	}

	for testNumber, test := range tests {

		actual := scalePoint(test.Scalar, test.Point)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v", expected)
			t.Logf("ACTUAL:   %v", actual)
			t.Logf("SCALAR = %d", test.Scalar)
			t.Logf("POINT:    %v", test.Point)
			continue
		}

	}
}
