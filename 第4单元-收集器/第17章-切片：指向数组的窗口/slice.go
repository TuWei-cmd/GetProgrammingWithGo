package main

import "fmt"

func main() {
	var planets = [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Sturn",
		"Uranus",
		"Neptune",
	}
	terrestrail := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]
	colonized := terrestrail[2:]
	fmt.Println(terrestrail, gasGiants, iceGiants, colonized)
	fmt.Println(gasGiants[0])

	giants := planets[4:]
	gas := giants[:2]
	ice := giants[2:4]
	fmt.Println(giants, gas, ice)
}
