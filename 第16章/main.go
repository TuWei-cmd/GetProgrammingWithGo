package main

import (
	"fmt"
)

type planet string

func main() {
	var planets = [...]planet{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Sturn",
		"Uranus",
		"Neptune",
	}
	// for i, dwarf := range planets {
	// 	fmt.Println(i, dwarf)
	// }
	planetsMarkII := planets
	planets[2] = "whoops"
	fmt.Println(planets)
	fmt.Println(planetsMarkII)
}
