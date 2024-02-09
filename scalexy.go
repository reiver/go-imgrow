package imgrow

func scaleXY(scalar int, x int, y int) (int, int) {
	x2 := x * scalar
	y2 := y * scalar

	return x2, y2
}
