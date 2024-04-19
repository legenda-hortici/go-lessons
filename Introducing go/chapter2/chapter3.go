package main

import "fmt"

func main() {
	fmt.Println("enter F: ")
	var f float64
	fmt.Scanf("%f", &f)

	m := f * 0.3048
	fmt.Println("M =", m)
}
