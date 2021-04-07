package main

import (
	"fmt"
	"math/rand"
)

func add(m int) float64 {
	var x float64
	switch m {
	case 1:
		x = 0.05
	case 2:
		x = 0.10
	case 3:
		x = 0.25
	}
	return x
}

func main() {
	var (
		store float64
	)
	for store < 20.0 {
		store += add(rand.Intn(3) + 1)
		fmt.Println(store)
	}
}
