package main 

import "fmt"

func max(xs ...int) int {
	max := 0
	for _, v := range xs {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	xs := []int{1,3,5,2,4}
	fmt.Println(max(xs...))
}