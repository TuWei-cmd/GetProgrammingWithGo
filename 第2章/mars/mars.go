package main

import (
	"fmt"
)

func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(130 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(18 *365 / 687)
	fmt.Print(" years old.\n")
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94.212417899808)
}