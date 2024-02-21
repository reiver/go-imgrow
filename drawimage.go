package imgrow

import (
	"image"
	"image/color"
	"image/draw"
)

// DrawImage returns an image that is 'img' where (only) drawing is scalad by 'scalar'.
// The actual image stays the same size.
// Only the drawing is scaled.
//
// So, for example:
//
//	var newImage draw.Image = imgrow.Image(2, oldImage)
//
// ... would return a new image that is just like 'oldImage' except any drawing on it is 2 times as big.
//
// Note that if you are trying to try a complex object and want the top-left corner to be at the original (unscaled) (x,y)
// rather (rather than the scaled position) then you need to move the top-left corner to:
//
//	xAdjusted := x / scalar
//	yAdjusted := y / scalar
func DrawImage(scalar int, img draw.Image) draw.Image {
	return internalDrawImage {
		img:img,
		scalar:scalar,
	}
}

var _ draw.Image = &internalDrawImage{}

type internalDrawImage struct {
	img draw.Image
	scalar int
}

func (receiver internalDrawImage) At(x, y int) color.Color {
	var img image.Image = receiver.img

	if nil == img {
		return nil
	}

	return img.At(x, y)
}

func (receiver internalDrawImage) Bounds() image.Rectangle {
	var img image.Image = receiver.img

	if nil == img {
		return image.Rectangle{}
	}

	return img.Bounds()
}

func (receiver internalDrawImage) ColorModel() color.Model {
	var img image.Image = receiver.img

	if nil == img {
		return nil
	}

	return img.ColorModel()
}

func (receiver internalDrawImage) Set(x int, y int, c color.Color) {
	var img draw.Image = receiver.img

	if nil == img {
		return
	}

	var scalar int = receiver.scalar

	xScaled, yScaled := scaleXY(scalar, x, y)

	for yi := yScaled; yi < (yScaled + scalar); yi++ {
		for xi := xScaled;  xi < (xScaled + scalar); xi++ {
			img.Set(xi, yi, c)
		}
	}
}
