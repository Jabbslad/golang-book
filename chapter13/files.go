package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
		return
	}
	
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		panic(err)
		return
	}
	str := string(bs)
	fmt.Println(str)
}