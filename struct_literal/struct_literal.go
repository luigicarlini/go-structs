package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
}

func main() {
	// composite literal; struct literal
	p1 := person{"Todd", "McLeod"}
	p2 := person{"Nina", "Simone"}
	fmt.Println(p1)
	fmt.Println(p2)

	a1 := struct {
		breed string
		name  string
	}{
		"German Sheperd",
		"Shasta",
	}
	fmt.Println(a1)
	m := meaning()
	fmt.Println(m)

}

// func (receiver) identifier(parameters) (returns) { <code> }
func meaning() int {
	return 42
}
