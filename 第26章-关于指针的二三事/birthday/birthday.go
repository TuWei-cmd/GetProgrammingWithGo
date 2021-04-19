package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func main() {
	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        17,
	}

	birthday(&rebecca)

	fmt.Printf("%+v\n", rebecca)
}
