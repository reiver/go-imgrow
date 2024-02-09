package imgrow

import (
	"image"
)

func scalePoint(scalar int, point image.Point) image.Point {
	x2, y2 := scaleXY(scalar, point.X, point.Y)

	return image.Point {
		X: x2,
		Y: y2,
	}
}
