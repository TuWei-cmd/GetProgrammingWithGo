package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\"%c", c.d, c.m, c.s, c.h)
}

type location struct {
	lat, long coordinate
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func main() {
	elysium := location{
		lat:  coordinate{4, 30, 0.0, 'N'},
		long: coordinate{135, 54, 0.0, 'E'},
	}
	fmt.Println("Elysium Planitia is at", elysium)
}
