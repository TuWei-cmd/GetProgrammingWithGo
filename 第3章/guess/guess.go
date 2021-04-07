package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	const myNumber = 14
	var num int
	for num = rand.Intn(100) + 1; num != myNumber;num = rand.Intn(100) + 1 {
		fmt.Printf("random number: %v which is too ", num)
		if num > myNumber {
			fmt.Printf("big!\n")
		} else{
			fmt.Printf("small!\n")
		}
	}
	fmt.Printf("Right! It's exactly %v.\n", num)
}