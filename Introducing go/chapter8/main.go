package main

import (
	"fmt"
	"golang/basic/shape"
)

func main() {
	circle := shape.NewCircle(5)
	rectangle := shape.NewRectangle(7)

	printPerimetr(circle)
	printPerimetr(rectangle)

}

func printPerimetr(s shape.Shape) {
	fmt.Println(s.Perimetr())
}
