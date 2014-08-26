package main

import "fmt"

func zero(x int) {
	x = 0
}

func zeroPtr(x *int) {
	*x = 0
}

func square(x *float64) {
	*x = *x * *x
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)

	zeroPtr(&x)
	fmt.Println(x)

	y := new(int)
	zeroPtr(y)
	fmt.Println(*y)

	z := 1.5
	square(&z)
	fmt.Println(z)
}