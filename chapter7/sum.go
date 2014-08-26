package main

import "fmt"

//explicit sum
func sum(xs []int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}

//recursive sum
func rsum(xs []int) int {
	if len(xs) == 0 {
		return 0
	}

	return xs[0] + rsum(xs[1:])
}

func main() {
	xs := []int{1,2,3,4,5}
	fmt.Println(sum(xs))
	fmt.Println(rsum(xs))
}