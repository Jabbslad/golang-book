package main

import "fmt"

func main() {
	x := [5]float64{ 
		98, 
		93, 
		77, 
		82, 
		//83 ,
	}

	fmt.Println(x)

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	slice1[0] = 9
	fmt.Println(slice1, slice2)

	slice2 = make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	arr := [5]int{1,2,3,4,5}
	slice3 := arr[0:3]
	slice4 := arr[0:5]
	fmt.Println(slice3, slice4)

	y := make(map[string]int)
	y["key"] = 10
	y["key2"] = 11
	fmt.Println(y)
	delete(y, "key2")
	fmt.Println(y)

	if name, ok := y["key3"]; ok {
		fmt.Println(name, ok)
	}

	mapOfMap := map[string]map[int]string {
		"first" : map[int]string {
			1 : "one",
		},
		"second" : map[int]string {
			2 : "two",
		},
	}

	if v, ok := mapOfMap["second"]; ok {
		fmt.Println(v[2])
	}

	z := []int{
    	48,96,86,68,
    	57,82,63,70,
    	37,34,83,27,
    	19,97, 9,17,
    }
    fmt.Println(smallest(z))

}

func smallest(arr []int) int {
	n := arr[0]
	for _, v := range arr {
		if v < n {
			n = v
		}
	}
	return n
}