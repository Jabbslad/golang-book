package main

import "testing"

func TestDistance(t *testing.T) {
	const x1, x2, y1, y2, out float64 = 2.0, 4.0, 2.0, 4.0, 2.8284271247461903
	if n := distance(x1, y1, x2, y2) ; n != out {
		t.Errorf("distance = %v, want %v", n, out)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{2.0, 2.0, 4.0, 4.0}
	out := 4.0
	if area := rectangle.area() ; area != out {
		t.Errorf("area = %v, want %v", area, 4)
	}
}