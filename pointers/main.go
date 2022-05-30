package main

import "fmt"

// ah whatever, these are the same as C

func main() {
	i := 1

	fmt.Println(i)
	fmt.Println(&i)
	test(&i)
	fmt.Println(i)
	fmt.Println(&i)
}

func test(num *int) {
	*num = 5
}
