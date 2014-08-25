package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	i := 0
	switch i {
		case 0: fmt.Println("Zero")
		default: fmt.Println("Unknown")
	}
}