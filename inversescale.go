package imgrow

// inverseScale basically does:
//
//	return n / scalar
//
// ... except, that it deals with the case where 0 == scalar in a special way.
func inverseScale(scalar int, n int) int {
	switch {
	case 0 == scalar:
		return divideByZero(n)
	default:
		return n / scalar
	}
}
