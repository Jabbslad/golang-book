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

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{y: 1, x: 2, r: 3}
	d := new(Circle)
	fmt.Println(c)
	fmt.Println("area :", d.area())
	fmt.Println(*d)
}