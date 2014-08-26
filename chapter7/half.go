package main 

import (
	"fmt"
)

func half(n int) (int, bool) {
	x := n / 2
	return x, n % 2 == 0
}

func main() {
	fmt.Println(half(1))
}