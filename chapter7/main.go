package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
	fmt.Println(*x)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	// x := 1.5
	// square(&x)
	x := 1
	y := 2
	swap(&x, &y)
}
