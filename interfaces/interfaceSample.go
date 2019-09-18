package main

import (
	"fmt"
)

type Shape interface { 
	Area() float64
	Volume() float64
}

type Square struct {
	length float64
}

func (sq Square) Area() float64 {
	return sq.length * sq.length
}

type Circle struct {
	radius float64
}

func (cr Circle) Area() float64 {
	return 3.14 * cr.radius
}

func main() {
	var s Shape
	fmt.Println(s)
	sq := Square{5}

	fmt.Printf("Value of square area type: %T, value %v\n", sq.Area(), sq.Area())

	cr := Circle{5}
	fmt.Printf("Value of Circle area type: %T, value %v\n", cr.Area(), cr.Area())

}
