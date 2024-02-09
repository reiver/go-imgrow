# go-imgrow

Package **imgrow** provides tools for **growing** and **image** by an integer value.

For example — if you have an image, then you make it 2× bigger, or 3× bigger, or 4× bigger, etc.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-imgrow

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-imgrow?status.svg)](https://godoc.org/sourcecode.social/reiver/go-imgrow)

## Example:

```go
import "sourcecode.social/reiver/go-imgrow"

// ...

var originalImage image.Image = ...

// ...

var scalar int = 400 // change this value to how many times bigger you want the image to be. here it is 400× bigger.

var newImage image.Image = imgrow(scalar, originalImage)

```

## Import

To import package **imgrow** use `import` code like the follownig:
```
import "sourcecode.social/reiver/go-imgrow"
```

## Installation

To install package **imgrow** do the following:
```
GOPROXY=direct go get https://sourcecode.social/reiver/go-imgrow
```

## Author

Package **imgrow** was written by [Charles Iliya Krempeaux](http://changelog.ca)
