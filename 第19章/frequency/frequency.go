package main

import "fmt"

func main() {
	temperature := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	frequency := make(map[float64]int, 6)

	for _, t := range temperature {
		frequency[t]++
	}
	fmt.Println(frequency)
}
