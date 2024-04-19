package shape

import (
	"math"
)

type Shape interface {
	Perimetr() float32
}

type Circle struct {
	radius float32
}

func NewCircle(radius float32) Circle {
	return Circle{
		radius: radius,
	}
}

func (c Circle) Perimetr() float32 {
	return 2 * c.radius * math.Pi
}

type Rectangle struct {
	rect float32
}

func (r Rectangle) Perimetr() float32 {
	return (r.rect + r.rect) * 2
}

func NewRectangle(rectangle float32) Rectangle {
	return Rectangle{
		rect: rectangle,
	}
}
