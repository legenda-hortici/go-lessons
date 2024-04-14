package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Perimetr() float32
}

type Circle struct {
	radius float32
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

func main() {
	circle := Circle{5}
	rectangle := Rectangle{7}

	printPerimetr(circle)
	printPerimetr(rectangle)

}

func printPerimetr(shape Shape) {
	fmt.Println(shape.Perimetr())
}
