package main

import "fmt"

type Shape interface { 
	Area() float64
	Volume() float64
}

type Cylinder struct {
	length float64
	breath float64
	height float64
}

type Cube struct {
	side float64
}

func (cyl Cylinder) Area() float64 {
	return 2 * 3.14 * cyl.breath * cyl.breath
}

func (cyl Cylinder) Volume() float64 {
	return 3.14 * cyl.breath * cyl.breath * cyl.height
}

func (cbe Cube) Area() float64 {
	return 3.14 * cbe.side
}

func (cbe Cube) Volume() float64 {
	return cbe.side * cbe.side * cbe.side
}

func main() {
	var s Shape = Cylinder{6, 3, 5}
	fmt.Printf("Value of square area type: %T, value %v\n", s.Area(), s.Area())

	s = Cube{5}
	fmt.Printf("Value of Circle area type: %T, value %v\n", s.Area(), s.Area())
}