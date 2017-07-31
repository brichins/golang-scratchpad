package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() int
}

type Circle struct {
	Radius int
}

func (shape Circle) Area() int {
	return int(float64(shape.Radius*shape.Radius) * math.Pi)
}

type Square struct {
	Length int
}

func (shape Square) Area() int {
	return shape.Length * shape.Length
}

func main() {
	s := Square{5}
	fmt.Printf("Square Length: %d\n", s.Length)
	fmt.Printf("Square Area: %d\n", s.Area())

	c := Circle{5}
	fmt.Printf("Circle Radius: %d\n", c.Radius)
	fmt.Printf("Circle Area: %d\n", c.Area())

	for _, shape := range []Shape{s, c} {
		println(shape.Area())
	}

}
