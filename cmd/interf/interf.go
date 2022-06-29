package interf

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	radius float64
}

func (rec Rectangle) Perimeter() float64 {
	return 2 * (rec.Height + rec.Width)
}

func (rec Rectangle) Area() float64 {
	return rec.Height * rec.Width
}

func (cir Circle) Area() float64 {
	return math.Pi * cir.radius * cir.radius
}

func (cir Circle) Perimeter() float64 {
	return 2 * math.Pi * cir.radius
}
