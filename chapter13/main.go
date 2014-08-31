package main

import (
	"fmt"
	"strings"
	"bytes"
)

func main() {
	fmt.Println(
		strings.Contains("test", "es"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),
		strings.Index("test", "e"),
		strings.Join([]string{"a", "b"}, "-"),
		strings.Repeat("a", 5),
		strings.Replace("aaaa", "a", "b", 2),
		strings.Split("a-b-c-d-e", "-"),
		strings.ToLower("TEST"),
		strings.ToUpper("test"),
	)

	arr := []byte("test")
	fmt.Println(arr)
	
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(str)

	var buf bytes.Buffer
	buf.Write([]byte("test"))
	fmt.Println("buffer:", buf)
	fmt.Println("bytes from buffer:", buf.Bytes())
	fmt.Println("bytes to string:", string(buf.Bytes()))

}