package main

import "fmt"

// Create 2 custom struct types called 'triangle' and 'square'
// The 'square' type should be a struct with a field called 'sideLength' of type float64
// The 'triangle' type should be a struct with a field called 'height' of type float64 and a field type 'base'
// Both types should have function called 'getArea' that returns the calculated area of the square or triangle
// Both types should have function called 'getArea' that returns the calculated area of the square or triangle
// Area of a triangle = 0.5*base*height
// Area of a square = sideLength * sideLength
// Add a 'shape' interface that defines a function called 'printArea'.
// This function should calculate the area of the given shape of and print it out to the terminal Design
// the interface so that 'printArea' function can be called with either a triangle or a Square
type shape interface {
	getArea() float64
}
type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	sq := square{
		sideLength: 5,
	}
	tri := triangle{
		height: 5,
		base:   6,
	}
	printArea(sq)
	printArea(tri)
}
