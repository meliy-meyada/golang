package main

import (
	"fmt"
)

type triangle struct {
	height,base float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}
func main() {
	fmt.Println("Interfaces Assignment")
	
	t := triangle{height: 2, base: 2}
	s := square{sideLength: 2}

	printArea(t)
	printArea(s)
}


