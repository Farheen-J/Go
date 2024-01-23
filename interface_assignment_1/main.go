package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type square struct {
	side float64
}

type triangle struct {
	base   float64
	height float64
}

func (s square) getArea() float64 {
	return math.Pow(s.side, 2)
}

func (t triangle) getArea() float64 {
	return (0.5 * t.base * t.height)
}

// Encapsulates the types
func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	s := square{10}
	printArea(s)
	t := triangle{base: 10, height: 10}
	printArea(t)
}
