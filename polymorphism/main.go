package main

import "fmt"

type Tickable interface {
	tick()
}
type Entity struct {
	Tickable
	name string
}

func (e *Entity) tick() {
	fmt.Println(e.name)
}

func main() {
	e := &Entity{
		name: "asd",
	}

	t := []Tickable{e}

	for _, t := range t {
		t.tick()
	}
}
