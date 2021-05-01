package main

import "fmt"

type Planets []string

func (p Planets) Newen() {
	for i := range p {
		p[i] = "New" + p[i]
	}
}

func main() {
	var planets = Planets{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Sturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(planets)
	planets[3:4].Newen()
	planets[6:8].Newen()
	fmt.Println(planets)
}
