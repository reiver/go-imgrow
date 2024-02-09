package imgrow

import (
	"testing"
)

func TestScaleXY(t *testing.T) {

	tests := []struct{
		Scalar int
		X int
		Y int
		ExpectedX int
		ExpectedY int
	}{
		{
			Scalar: -2,
			X: 5,
			Y: 7,
			ExpectedX: -10,
			ExpectedY: -14,
		},
		{
			Scalar: -1,
			X: 5,
			Y: 7,
			ExpectedX: -5,
			ExpectedY: -7,
		},
		{
			Scalar: 0,
			X: 5,
			Y: 7,
			ExpectedX: 0,
			ExpectedY: 0,
		},
		{
			Scalar: 1,
			X: 5,
			Y: 7,
			ExpectedX: 5,
			ExpectedY: 7,
		},
		{
			Scalar: 2,
			X: 5,
			Y: 7,
			ExpectedX: 10,
			ExpectedY: 14,
		},
		{
			Scalar: 3,
			X: 5,
			Y: 7,
			ExpectedX: 15,
			ExpectedY: 21,
		},




		{
			Scalar: 100,
			X: -11,
			Y: 13,
			ExpectedX: -1100,
			ExpectedY: 1300,
		},
	}

	for testNumber, test := range tests {

		actualX, actualY := scaleXY(test.Scalar, test.X, test.Y)
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
