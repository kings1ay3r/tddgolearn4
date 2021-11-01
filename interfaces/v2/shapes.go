package main

import "math"

type Rectangle struct {
	Width  float64
	height float64
}
type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.height + rectangle.Width)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.height
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}
