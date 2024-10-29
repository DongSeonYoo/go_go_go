package structs

import "math"

// The Shape interface has Area
type Shape interface {
	Area() float64
}

// The Struct of Triangle
type Triangle struct {
	Base   float64
	Height float64
}

// The Struct of Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// The Struct of Circle
type Circle struct {
	radius float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Height)
}

// fomular of Rectangle's Area: Height * Width
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// fomular of Circle's Area: PIE * (Raduis)^
func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}

// fomular of Triangle's Area: (Base * Height) / 2
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
