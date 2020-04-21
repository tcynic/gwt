package main

import "math"

// Shape is an interface for float64
type Shape interface {
	Area() float64
}

// Rectangle is two floating point values Width and Height
type Rectangle struct {
	Width  float64
	Height float64
}

// Area is a method that takes a height and width and returns float64
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle takes a float64 value
type Circle struct {
	Radius float64
}

// Area returns the value of the 2*radius*pi
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle takes 2 float64s, base and height
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the value of the base*heigh*.5
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter takes a struct Rectangle and returns its Perimeter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
