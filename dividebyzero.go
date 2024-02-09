package imgrow

func divideByZero(n int) int {
	switch {
	case 0 < n:
		return intMax
	case n < 0:
		return intMin
	default: // 0 == n
		return 0
	}
}
