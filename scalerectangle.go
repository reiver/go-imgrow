package imgrow

import (
	"image"
)

func scaleRectangle(scalar int, rectangle image.Rectangle) image.Rectangle {
	return image.Rectangle{
		Min: scalePoint(scalar, rectangle.Min),
		Max: scalePoint(scalar, rectangle.Max),
	}
}
