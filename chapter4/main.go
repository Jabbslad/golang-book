package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	fmt.Println(FtoC(input))
}

func FtoC(f float64) float64 {
	return (f - 32) * 5/9
}