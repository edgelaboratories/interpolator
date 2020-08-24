# Interpolator

[![GoDoc](https://godoc.org/github.com/edgelaboratories/interpolator?status.png)](http://godoc.org/github.com/edgelaboratories/interpolator)
![Build Status](https://github.com/edgelaboratories/interpolator/workflows/Test/badge.svg)
![GolangCI Lint](https://github.com/edgelaboratories/interpolator/workflows/GolangCI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/edgelaboratories/interpolator)](https://goreportcard.com/report/github.com/edgelaboratories/interpolator)

## Description

Package `interpolator` provides univariate data interpolators:

* [piecewise-constant](piecewise_constant.go) interpolator
* [piecewise-linear](piecewise_linear.go) interpolator
* [piecewise-linear with threshold](piecewise_linear_threshold.go): the interpolated value is truncated to the closest in the data range, when the input point is out of the data domain, in order to prevent extrapolation effects
* [piecewise-geometric](geometric.go)
* [piecewise-geometric on square-root factor](geometric_sqrt.go): the interpolated value depends on the square root of the normalized distance from data points

The input data is specified by means of a nonempty [slice of two-dimensional points](xy.go) `XYs`. If a single data point is provided, the resulting interpolator **treats the input as a constant** for all abscissae.

## Installation

    go get -u github.com/edgelaboratories/interpolator

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/edgelaboratories/interpolator"
)

func main() {
	xys := interpolator.XYs{
		{
			X: 0.0,
			Y: 1.2,
		},
		{
			X: 0.5,
			Y: 1.0,
		},
		{
			X: 1.0,
			Y: 1.4,
		},
	}
	interp, err := interpolator.NewGeometric(xys)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("value at 0.75 is %0.2f\n", interp.Value(0.75))
	fmt.Printf("gradient at 0.75 is %0.2f\n", interp.Gradient(0.75))
}
```
