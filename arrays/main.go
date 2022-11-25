package main

import "fmt"

func main() {
	// Arrays
	a := [5]string{"A", "B", "C", "D", "E"}
	b := [...]string{1: "B", 3: "D"} // Index:Value ile belirli bir indexsteki veriyi initialize edebiliyoruz

	// Arrays are so cool!
	// I wish we had this in Java too
	fmt.Println(a[:3]) // prints out [A B C]
	fmt.Println(b[:3]) // prints out [A B C]

	c := a[2:4]
	fmt.Println(cap(c))
	fmt.Println(len(c))
	fmt.Println(c)
}
