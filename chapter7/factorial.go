package main

import (
	"fmt"
	"os"
	"strconv"
)

func factorial(x int64) int64 {
	if x==0 {
		return 1
	}

	return x * factorial(x - 1)
}

func main() {
	x, _ := strconv.ParseInt(os.Args[1], 10, 0)
	fmt.Println(factorial(x))
}