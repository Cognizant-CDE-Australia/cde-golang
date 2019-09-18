package main

import "fmt"

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Cylinder struct {
	length float64
	breath float64
	height float64
}

func (cyl Cylinder) Area() float64 {
	return 2 * 3.14 * cyl.breath * cyl.breath
}

func (cyl Cylinder) Volume() float64 {
	return 3.14 * cyl.breath * cyl.breath * cyl.height
}

func main() {
	c := Cylinder{6, 3, 5}
	var s Shape = c
	var o Object = c
	fmt.Println("volume of s of interface type Shape is", s.Area())
	fmt.Println("area of o of interface type Object is", o.Volume())
}
