package main

import (
	"fmt"
	"time"
)

func sleep(n time.Duration) {
	select {
		case <- time.After(time.Second * n):
	}
}

func main() {
	fmt.Println(time.Now())
	sleep(10)
	fmt.Println(time.Now())
}