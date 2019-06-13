# Interpolator

[![GoDoc](https://godoc.org/github.com/edgelaboratories/interpolator?status.png)](http://godoc.org/github.com/edgelaboratories/interpolator)
[![Build Status](https://api.travis-ci.org/edgelaboratories/interpolator.svg?branch=master)](https://travis-ci.org/edgelaboratories/interpolator)
[![Go Report Card](https://goreportcard.com/badge/github.com/edgelaboratories/interpolator)](https://goreportcard.com/report/github.com/edgelaboratories/interpolator)

## Description

Package `interpolator` provides univariate data interpolators.

## Installation

    go get -u github.com/edgelaboratories/interpolator

## Example

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	xys := XYs{
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
	interp, err := NewGeometric(xys)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("value at 0.75 is %0.2f\n", interp.Value(0.75))
	fmt.Printf("gradient at 0.75 is %0.2f\n", interp.Gradient(0.75))
}
```