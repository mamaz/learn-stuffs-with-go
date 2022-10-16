package main

import (
	"fmt"
)

type Leg struct {
	ID string
}

func (l *Leg) Walk() {
	fmt.Println("walk")
}

type Person struct {
	Leg
}

func main() {
	l := Leg{ID: "12"}
	p := Person{}
	p.Leg = l

	fmt.Println(p.ID)
	p.Walk()
}
