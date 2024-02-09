package imgrow

import (
	"testing"

	"image"
)

func TestScaleRectangle(t *testing.T) {

	tests := []struct{
		Scalar int
		Rectangle image.Rectangle
		Expected image.Rectangle
	}{
		{
			Scalar: -2,
			Rectangle: image.Rectangle{
				Min: image.Point{
					X: 2,
					Y: 3,
				},
				Max: image.Point{
					X: 5,
					Y: 7,
				},
			},
			Expected: image.Rectangle{
				Min: image.Point{
					X: -4,
					Y: -6,
				},
				Max: image.Point{
					X: -10,
					Y: -14,
				},
			},
		},
		{
			Scalar: -1,
			Rectangle: image.Rectangle{
				Min: image.Point{
					X: 0,
					Y: 0,
				},
				Max: image.Point{
					X: 5,
					Y: 7,
				},
			},
			Expected: image.Rectangle{
				Min: image.Point{
					X: 0,
					Y: 0,
				},
				Max: image.Point{
					X: -5,
					Y: -7,
				},
			},
		},
		{
			Scalar: 0,
			Rectangle: image.Rectangle{
				Max: image.Point{
					X: 5,
					Y: 7,
				},
			},
			Expected: image.Rectangle{
				Max: image.Point{
					X: 0,
					Y: 0,
				},
			},
		},
		{
			Scalar: 1,
			Rectangle: image.Rectangle{
				Min: image.Point{
					X: 0,
					Y: 0,
				},
				Max: image.Point{
					X: 5,
					Y: 7,
				},
			},
			Expected: image.Rectangle{
				Min: image.Point{
					X: 0,
					Y: 0,
				},
				Max: image.Point{
					X: 5,
					Y: 7,
				},
			},
		},
		{
			Scalar: 2,
			Rectangle: image.Rectangle{
				Min: image.Point{
					X: 2,
					Y: 3,
				},
				Max: image.Point{
					X: 5,
					Y: 7,
				},
			},
			Expected: image.Rectangle{
				Min: image.Point{
					X: 4,
					Y: 6,
				},
				Max: image.Point{
					X: 10,
					Y: 14,
				},
			},
		},
		{
			Scalar: 3,
			Rectangle: image.Rectangle{
				Min: image.Point{
					X: 1,
					Y: 2,
				},
				Max: image.Point{
					X: 5,
					Y: 7,
				},
			},
			Expected: image.Rectangle{
				Min: image.Point{
					X: 3,
					Y: 6,
				},
				Max: image.Point{
					X: 15,
					Y: 21,
				},
			},
		},




		{
			Scalar: 100,
			Rectangle: image.Rectangle{
				Min: image.Point{
					X: -3,
					Y: 0,
				},
				Max: image.Point{
					X: -11,
					Y: 13,
				},
			},
			Expected: image.Rectangle{
				Min: image.Point{
					X: -300,
					Y: 0,
				},
				Max: image.Point{
					X: -1100,
					Y: 1300,
				},
			},
		},
	}

	for testNumber, test := range tests {

		actual := scaleRectangle(test.Scalar, test.Rectangle)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v", expected)
			t.Logf("ACTUAL:   %v", actual)
			t.Logf("SCALAR = %d", test.Scalar)
			t.Logf("POINT:    %v", test.Rectangle)
			continue
		}

	}
}
