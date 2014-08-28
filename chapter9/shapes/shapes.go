package main

import (
	"math"
	"fmt"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Circle struct {
	x, y, r float64
}

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (m *MultiShape) area() float64 {
	total := 0.0
	for _, v := range m.shapes {
		total += v.area()
	}
	return total
}

func main() {
	r1 := Rectangle{2, 2, 4, 4}
	m := MultiShape{[]Shape{&r1}}
	fmt.Println(m.area())
}