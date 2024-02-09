package imgrow

func inverseScaleXY(scalar int, x int, y int) (int, int) {
	return inverseScale(scalar, x), inverseScale(scalar, y)
}
