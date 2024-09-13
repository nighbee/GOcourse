package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
func PrintArea(s Shape) {
	fmt.Println(s.area())
}
func main() {
	c := Circle{radius: 5}
	r := Rectangle{width: 10, height: 5}
	PrintArea(c)
	PrintArea(r)
}
