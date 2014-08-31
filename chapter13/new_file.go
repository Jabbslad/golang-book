package main

import "os"

func main() {
	file, err := os.Create("test_create.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("Eyo!")
}