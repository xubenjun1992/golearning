package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width float64
	Hight float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Hight * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Hight + r.Width)
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	var r = Rectangle{
		Width: 10.2,
		Hight: 2.5,
	}
	var c = Circle{
		Radius: 1.5,
	}
	fmt.Printf("长方形面积: %.2f\n长方形周长: %.2f\n", r.Area(), r.Perimeter())
	fmt.Printf("圆形面积: %.2f\n圆形周长: %.2f\n", c.Area(), c.Perimeter())
}
