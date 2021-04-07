package main 

import (
	"fmt"
	"math/big"
)

func main() {
	distance := new(big.Int)
	distance.SetString("236000000000000000", 10)
	lightSpeed := big.NewInt(299792)
	secondsPreYear := big.NewInt(3600*24*365)
	seconds := new(big.Int)
	years := new(big.Int)
	seconds.Div(distance, lightSpeed)
	years.Div(seconds, secondsPreYear)
	fmt.Println("Canis Major Dwarf is", distance, "km away.")
	fmt.Println("That is", years, "years of travel at light speed.")
}
