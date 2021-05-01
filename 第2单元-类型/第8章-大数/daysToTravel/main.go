package main

import (
	"fmt"
	"math/big"
)

func alpha() {
	fmt.Println("飞行至半人马座阿尔法星所需的天数：")

	const lightSpeed = 299792 // km/s
	const secondsPerDay = 86400

	var distance int64 = 41.3e12 //km
	fmt.Println("Alpha Centauri is", distance, "km away.")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.")
}

func andromeda() {
	fmt.Println("飞行至仙女星系所需的天数：")
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance, "km away.")
	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)
	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println("That is", days, "days of travel at light speed.")
}

func main() {
	fmt.Println("   ")
	alpha()
	andromeda()
	fmt.Println("   ")
}