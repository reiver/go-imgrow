package imgrow

import (
	"testing"
)

func TestInverseScaleXY(t *testing.T) {

	tests := []struct{
		Scalar int
		X int
		Y int
		ExpectedX int
		ExpectedY int
	}{
		{
			Scalar: 0,
			X: 0,
			Y: 0,
			ExpectedX: 0,
			ExpectedY: 0,
		},
		{
			Scalar: 0,
			X: 1,
			Y: 0,
			ExpectedX: intMax,
			ExpectedY: 0,
		},
		{
			Scalar: 0,
			X: 0,
			Y: 1,
			ExpectedX: 0,
			ExpectedY: intMax,
		},
		{
			Scalar: 0,
			X: -1,
			Y: 0,
			ExpectedX: intMin,
			ExpectedY: 0,
		},
		{
			Scalar: 0,
			X: 0,
			Y: -1,
			ExpectedX: 0,
			ExpectedY: intMin,
		},
		{
			Scalar: 0,
			X: 1,
			Y: 1,
			ExpectedX: intMax,
			ExpectedY: intMax,
		},
		{
			Scalar: 0,
			X: -1,
			Y: -1,
			ExpectedX: intMin,
			ExpectedY: intMin,
		},
		{
			Scalar: 0,
			X: -1,
			Y: 1,
			ExpectedX: intMin,
			ExpectedY: intMax,
		},
		{
			Scalar: 0,
			X: 1,
			Y: -1,
			ExpectedX: intMax,
			ExpectedY: intMin,
		},
		{
			Scalar: 0,
			X: 5,
			Y: -7,
			ExpectedX: intMax,
			ExpectedY: intMin,
		},









		{
			Scalar: 1,
			X: 0,
			Y: 0,
			ExpectedX: 0,
			ExpectedY: 0,
		},
		{
			Scalar: 1,
			X: 1,
			Y: 0,
			ExpectedX: 1,
			ExpectedY: 0,
		},
		{
			Scalar: 1,
			X: 2,
			Y: 0,
			ExpectedX: 2,
			ExpectedY: 0,
		},
		{
			Scalar: 1,
			X: 3,
			Y: 0,
			ExpectedX: 3,
			ExpectedY: 0,
		},
		{
			Scalar: 1,
			X: 4,
			Y: 0,
			ExpectedX: 4,
			ExpectedY: 0,
		},
		{
			Scalar: 1,
			X: 5,
			Y: 0,
			ExpectedX: 5,
			ExpectedY: 0,
		},
		{
			Scalar: 1,
			X: 6,
			Y: 0,
			ExpectedX: 6,
			ExpectedY: 0,
		},
		{
			Scalar: 1,
			X: 7,
			Y: 0,
			ExpectedX: 7,
			ExpectedY: 0,
		},
		{
			Scalar: 1,
			X: 8,
			Y: 0,
			ExpectedX: 8,
			ExpectedY: 0,
		},
		{
			Scalar: 1,
			X: 9,
			Y: 0,
			ExpectedX: 9,
			ExpectedY: 0,
		},
		{
			Scalar: 1,
			X: 10,
			Y: 0,
			ExpectedX: 10,
			ExpectedY: 0,
		},

		{
			Scalar: 1,
			X: 0,
			Y: 1,
			ExpectedX: 0,
			ExpectedY: 1,
		},
		{
			Scalar: 1,
			X: 1,
			Y: 1,
			ExpectedX: 1,
			ExpectedY: 1,
		},
		{
			Scalar: 1,
			X: 2,
			Y: 1,
			ExpectedX: 2,
			ExpectedY: 1,
		},
		{
			Scalar: 1,
			X: 3,
			Y: 1,
			ExpectedX: 3,
			ExpectedY: 1,
		},
		{
			Scalar: 1,
			X: 4,
			Y: 1,
			ExpectedX: 4,
			ExpectedY: 1,
		},
		{
			Scalar: 1,
			X: 5,
			Y: 1,
			ExpectedX: 5,
			ExpectedY: 1,
		},
		{
			Scalar: 1,
			X: 6,
			Y: 1,
			ExpectedX: 6,
			ExpectedY: 1,
		},
		{
			Scalar: 1,
			X: 7,
			Y: 1,
			ExpectedX: 7,
			ExpectedY: 1,
		},
		{
			Scalar: 1,
			X: 8,
			Y: 1,
			ExpectedX: 8,
			ExpectedY: 1,
		},
		{
			Scalar: 1,
			X: 9,
			Y: 1,
			ExpectedX: 9,
			ExpectedY: 1,
		},
		{
			Scalar: 1,
			X: 10,
			Y: 1,
			ExpectedX: 10,
			ExpectedY: 1,
		},

		{
			Scalar: 1,
			X: 0,
			Y: 2,
			ExpectedX: 0,
			ExpectedY: 2,
		},
		{
			Scalar: 1,
			X: 1,
			Y: 2,
			ExpectedX: 1,
			ExpectedY: 2,
		},
		{
			Scalar: 1,
			X: 2,
			Y: 2,
			ExpectedX: 2,
			ExpectedY: 2,
		},
		{
			Scalar: 1,
			X: 3,
			Y: 2,
			ExpectedX: 3,
			ExpectedY: 2,
		},
		{
			Scalar: 1,
			X: 4,
			Y: 2,
			ExpectedX: 4,
			ExpectedY: 2,
		},
		{
			Scalar: 1,
			X: 5,
			Y: 2,
			ExpectedX: 5,
			ExpectedY: 2,
		},
		{
			Scalar: 1,
			X: 6,
			Y: 2,
			ExpectedX: 6,
			ExpectedY: 2,
		},
		{
			Scalar: 1,
			X: 7,
			Y: 2,
			ExpectedX: 7,
			ExpectedY: 2,
		},
		{
			Scalar: 1,
			X: 8,
			Y: 2,
			ExpectedX: 8,
			ExpectedY: 2,
		},
		{
			Scalar: 1,
			X: 9,
			Y: 2,
			ExpectedX: 9,
			ExpectedY: 2,
		},
		{
			Scalar: 1,
			X: 10,
			Y: 2,
			ExpectedX: 10,
			ExpectedY: 2,
		},

		{
			Scalar: 1,
			X: 0,
			Y: 3,
			ExpectedX: 0,
			ExpectedY: 3,
		},
		{
			Scalar: 1,
			X: 1,
			Y: 3,
			ExpectedX: 1,
			ExpectedY: 3,
		},
		{
			Scalar: 1,
			X: 2,
			Y: 3,
			ExpectedX: 2,
			ExpectedY: 3,
		},
		{
			Scalar: 1,
			X: 3,
			Y: 3,
			ExpectedX: 3,
			ExpectedY: 3,
		},
		{
			Scalar: 1,
			X: 4,
			Y: 3,
			ExpectedX: 4,
			ExpectedY: 3,
		},
		{
			Scalar: 1,
			X: 5,
			Y: 3,
			ExpectedX: 5,
			ExpectedY: 3,
		},
		{
			Scalar: 1,
			X: 6,
			Y: 3,
			ExpectedX: 6,
			ExpectedY: 3,
		},
		{
			Scalar: 1,
			X: 7,
			Y: 3,
			ExpectedX: 7,
			ExpectedY: 3,
		},
		{
			Scalar: 1,
			X: 8,
			Y: 3,
			ExpectedX: 8,
			ExpectedY: 3,
		},
		{
			Scalar: 1,
			X: 9,
			Y: 3,
			ExpectedX: 9,
			ExpectedY: 3,
		},
		{
			Scalar: 1,
			X: 10,
			Y: 3,
			ExpectedX: 10,
			ExpectedY: 3,
		},

		{
			Scalar: 1,
			X: 0,
			Y: 4,
			ExpectedX: 0,
			ExpectedY: 4,
		},
		{
			Scalar: 1,
			X: 1,
			Y: 4,
			ExpectedX: 1,
			ExpectedY: 4,
		},
		{
			Scalar: 1,
			X: 2,
			Y: 4,
			ExpectedX: 2,
			ExpectedY: 4,
		},
		{
			Scalar: 1,
			X: 3,
			Y: 4,
			ExpectedX: 3,
			ExpectedY: 4,
		},
		{
			Scalar: 1,
			X: 4,
			Y: 4,
			ExpectedX: 4,
			ExpectedY: 4,
		},
		{
			Scalar: 1,
			X: 5,
			Y: 4,
			ExpectedX: 5,
			ExpectedY: 4,
		},
		{
			Scalar: 1,
			X: 6,
			Y: 4,
			ExpectedX: 6,
			ExpectedY: 4,
		},
		{
			Scalar: 1,
			X: 7,
			Y: 4,
			ExpectedX: 7,
			ExpectedY: 4,
		},
		{
			Scalar: 1,
			X: 8,
			Y: 4,
			ExpectedX: 8,
			ExpectedY: 4,
		},
		{
			Scalar: 1,
			X: 9,
			Y: 4,
			ExpectedX: 9,
			ExpectedY: 4,
		},
		{
			Scalar: 1,
			X: 10,
			Y: 4,
			ExpectedX: 10,
			ExpectedY: 4,
		},

		{
			Scalar: 1,
			X: 0,
			Y: 5,
			ExpectedX: 0,
			ExpectedY: 5,
		},
		{
			Scalar: 1,
			X: 1,
			Y: 5,
			ExpectedX: 1,
			ExpectedY: 5,
		},
		{
			Scalar: 1,
			X: 2,
			Y: 5,
			ExpectedX: 2,
			ExpectedY: 5,
		},
		{
			Scalar: 1,
			X: 3,
			Y: 5,
			ExpectedX: 3,
			ExpectedY: 5,
		},
		{
			Scalar: 1,
			X: 4,
			Y: 5,
			ExpectedX: 4,
			ExpectedY: 5,
		},
		{
			Scalar: 1,
			X: 5,
			Y: 5,
			ExpectedX: 5,
			ExpectedY: 5,
		},
		{
			Scalar: 1,
			X: 6,
			Y: 5,
			ExpectedX: 6,
			ExpectedY: 5,
		},
		{
			Scalar: 1,
			X: 7,
			Y: 5,
			ExpectedX: 7,
			ExpectedY: 5,
		},
		{
			Scalar: 1,
			X: 8,
			Y: 5,
			ExpectedX: 8,
			ExpectedY: 5,
		},
		{
			Scalar: 1,
			X: 9,
			Y: 5,
			ExpectedX: 9,
			ExpectedY: 5,
		},
		{
			Scalar: 1,
			X: 10,
			Y: 5,
			ExpectedX: 10,
			ExpectedY: 5,
		},
	}

	for testNumber, test := range tests {

		actualX, actualY := inverseScaleXY(test.Scalar, test.X, test.Y)
		expectedX, expectedY := test.ExpectedX, test.ExpectedY

		if expectedX != actualX || expectedY != actualY {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED (X,Y) = (%d, %d)", expectedX, expectedY)
			t.Logf("ACTUAL (X,Y) = (%d, %d)", actualX, actualY)
			t.Logf("SCALAR = %d", test.Scalar)
			t.Logf("(X,Y) = (%d, %d)", test.X, test.Y)
			continue
		}

	}
}
