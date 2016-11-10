package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}
type Rectangle struct {
	width float64
	height float64
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Shape interface {
	 area() float64
}
type MultipleShape struct {
	shapes []Shape
}	

func main() {
	multiShape := MultipleShape {
		shapes: []Shape{
			Circle { 0, 0, 5},
			Rectangle{ 10, 10 },
		},
	}
	for name, v := range multiShape.shapes {
		fmt.Println(name, v.area())
	}
	fmt.Println()
}
