package main

import "fmt"

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// create Area() method for square
func (s square) Area() float64 {
	return s.length * s.length
}

// create Area() method for circle
func (c circle) Area() float64 {
	pi := 3.14
	return pi * float64(c.radius*c.radius)
}

// create a interface SHAPE, which defines anything as SHAPE that has Area() method
type shape interface {
	Area() float64
}

// create info() shape and prints area of that shape
func info(s shape) {
	result := s.Area()
	fmt.Println("area: ", result)
}

func main() {
	// create a square
	sq := square{
		length: 4,
	}

	// create a circle
	ci := circle{
		radius: 5,
	}
	// call info() to return the area of square and circle
	info(sq)
	info(ci)
}
