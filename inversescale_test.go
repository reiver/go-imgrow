package imgrow

import (
	"testing"
)

func TestInverseScale(t *testing.T) {

	tests := []struct{
		Scalar int
		N int
		Expected int
	}{
		{
			Scalar: 0,
			N: -8,
			Expected: intMin,
		},
		{
			Scalar: 0,
			N: -7,
			Expected: intMin,
		},
		{
			Scalar: 0,
			N: -6,
			Expected: intMin,
		},
		{
			Scalar: 0,
			N: -5,
			Expected: intMin,
		},
		{
			Scalar: 0,
			N: -4,
			Expected: intMin,
		},
		{
			Scalar: 0,
			N: -3,
			Expected: intMin,
		},
		{
			Scalar: 0,
			N: -2,
			Expected: intMin,
		},
		{
			Scalar: 0,
			N: -1,
			Expected: intMin,
		},
		{
			Scalar: 0,
			N: 0,
			Expected: 0,
		},
		{
			Scalar: 0,
			N: 1,
			Expected: intMax,
		},
		{
			Scalar: 0,
			N: 2,
			Expected: intMax,
		},
		{
			Scalar: 0,
			N: 3,
			Expected: intMax,
		},
		{
			Scalar: 0,
			N: 4,
			Expected: intMax,
		},
		{
			Scalar: 0,
			N: 5,
			Expected: intMax,
		},
		{
			Scalar: 0,
			N: 6,
			Expected: intMax,
		},
		{
			Scalar: 0,
			N: 7,
			Expected: intMax,
		},
		{
			Scalar: 0,
			N: 8,
			Expected: intMax,
		},



		{
			Scalar: 1,
			N: -8,
			Expected: -8,
		},
		{
			Scalar: 1,
			N: -7,
			Expected: -7,
		},
		{
			Scalar: 1,
			N: -6,
			Expected: -6,
		},
		{
			Scalar: 1,
			N: -5,
			Expected: -5,
		},
		{
			Scalar: 1,
			N: -4,
			Expected: -4,
		},
		{
			Scalar: 1,
			N: -3,
			Expected: -3,
		},
		{
			Scalar: 1,
			N: -2,
			Expected: -2,
		},
		{
			Scalar: 1,
			N: -1,
			Expected: -1,
		},
		{
			Scalar: 1,
			N: 0,
			Expected: 0,
		},
		{
			Scalar: 1,
			N: 1,
			Expected: 1,
		},
		{
			Scalar: 1,
			N: 2,
			Expected: 2,
		},
		{
			Scalar: 1,
			N: 3,
			Expected: 3,
		},
		{
			Scalar: 1,
			N: 4,
			Expected: 4,
		},
		{
			Scalar: 1,
			N: 5,
			Expected: 5,
		},
		{
			Scalar: 1,
			N: 6,
			Expected: 6,
		},
		{
			Scalar: 1,
			N: 7,
			Expected: 7,
		},
		{
			Scalar: 1,
			N: 8,
			Expected: 8,
		},



		{
			Scalar: 2,
			N: -8,
			Expected: -4,
		},
		{
			Scalar: 2,
			N: -7,
			Expected: -3,
		},
		{
			Scalar: 2,
			N: -6,
			Expected: -3,
		},
		{
			Scalar: 2,
			N: -5,
			Expected: -2,
		},
		{
			Scalar: 2,
			N: -4,
			Expected: -2,
		},
		{
			Scalar: 2,
			N: -3,
			Expected: -1,
		},
		{
			Scalar: 2,
			N: -2,
			Expected: -1,
		},
		{
			Scalar: 2,
			N: -1,
			Expected: 0,
		},
		{
			Scalar: 2,
			N: 0,
			Expected: 0,
		},
		{
			Scalar: 2,
			N: 1,
			Expected: 0,
		},
		{
			Scalar: 2,
			N: 2,
			Expected: 1,
		},
		{
			Scalar: 2,
			N: 3,
			Expected: 1,
		},
		{
			Scalar: 2,
			N: 4,
			Expected: 2,
		},
		{
			Scalar: 2,
			N: 5,
			Expected: 2,
		},
		{
			Scalar: 2,
			N: 6,
			Expected: 3,
		},
		{
			Scalar: 2,
			N: 7,
			Expected: 3,
		},
		{
			Scalar: 2,
			N: 8,
			Expected: 4,
		},



		{
			Scalar: 3,
			N: -8,
			Expected: -2,
		},
		{
			Scalar: 3,
			N: -7,
			Expected: -2,
		},
		{
			Scalar: 3,
			N: -6,
			Expected: -2,
		},
		{
			Scalar: 3,
			N: -5,
			Expected: -1,
		},
		{
			Scalar: 3,
			N: -4,
			Expected: -1,
		},
		{
			Scalar: 3,
			N: -3,
			Expected: -1,
		},
		{
			Scalar: 3,
			N: -2,
			Expected: 0,
		},
		{
			Scalar: 3,
			N: -1,
			Expected: 0,
		},
		{
			Scalar: 3,
			N: 0,
			Expected: 0,
		},
		{
			Scalar: 3,
			N: 1,
			Expected: 0,
		},
		{
			Scalar: 3,
			N: 2,
			Expected: 0,
		},
		{
			Scalar: 3,
			N: 3,
			Expected: 1,
		},
		{
			Scalar: 3,
			N: 4,
			Expected: 1,
		},
		{
			Scalar: 3,
			N: 5,
			Expected: 1,
		},
		{
			Scalar: 3,
			N: 6,
			Expected: 2,
		},
		{
			Scalar: 3,
			N: 7,
			Expected: 2,
		},
		{
			Scalar: 3,
			N: 8,
			Expected: 2,
		},
	}

	for testNumber, test := range tests {

		actual := inverseScale(test.Scalar, test.N)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("SCALAR: %d", test.Scalar)
			t.Logf("N:      %d", test.N)
			continue
		}
	}
}
