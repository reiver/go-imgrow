package imgrow

import (
	"image"
	"image/color"
)

// Image returns an image that is 'img' scalad by 'scalar'.
//
// So, for example:
//
//	var newImage image.Image = imgrow.Image(2, oldImage)
//
// ... would return a new image that is just like 'oldImage' except that it is 2 times as big.
func Image(scalar int, img image.Image) image.Image {
	return internalImage {
		img:img,
		scalar:scalar,
	}
}

var _ image.Image = &internalImage{}

type internalImage struct {
	img image.Image
	scalar int
}

func (receiver internalImage) At(x, y int) color.Color {
	var img image.Image = receiver.img

	if nil == img {
		return nil
	}

	var scalar int = receiver.scalar

	x2, y2 := inverseScaleXY(scalar, x, y)

	return img.At(x2, y2)
}

func (receiver internalImage) Bounds() image.Rectangle {
	var img image.Image = receiver.img

	if nil == img {
		return image.Rectangle{}
	}

	return scaleRectangle(receiver.scalar, img.Bounds())
}

func (receiver internalImage) ColorModel() color.Model {
	var img image.Image = receiver.img

	if nil == img {
		return nil
	}

	return img.ColorModel()
}
