package main

import (
	"fmt"
	"math"
)

type square struct {
	side_length float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return math.Pow(s.side_length, 2)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {

	switch s.(type) {
	case square:
		fmt.Printf("Side Length of Sqaure: %v Area of Square: %v\n"
		, s.(square).side_length, s.(square).area())
	case circle:
		fmt.Printf("Radius of Circle: %v Area of Circle: %v\n"
		, s.(circle).radius, s.(circle).area())
	}

}

func main() {
	s := square{
		side_length: 10,
	}
	info(s)

	c := circle{
		radius: 4,
	}
	info(c)
}
