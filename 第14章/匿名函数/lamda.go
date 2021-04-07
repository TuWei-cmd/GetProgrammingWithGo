package main

import (
	"fmt"
	"math/rand"
)

var f = func() {
	fmt.Println("Dress up for the masquerade.")
}

func main() {
	f()

	funFn := func(message string) {
		fmt.Println(message)
	}
	funFn("It is a fun funFn!!HHH")

	func() {
		fmt.Println("Functions anonymous")
	}()

	a := rand.Intn(124)
	fmt.Println(a)
}
