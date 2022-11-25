package main

import "fmt"

func main() {
	mymap := map[string]string{"ano": "babo"}

	fmt.Println(mymap)
	fmt.Println(mymap["ano"])
	mymap["Hello"] = "World"
	fmt.Println(mymap["Hello"])
}
