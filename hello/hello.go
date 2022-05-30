package main

import (
	"fmt"
)

func main() {
	a, b, c := Test()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func Test() (string, string, int32) {
	return "a", "b", 1
}
