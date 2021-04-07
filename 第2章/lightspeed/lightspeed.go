package main

import "fmt"

func main() {
	const speed = 	100800 // km/h
	var distance = 96300000 // km
	var time = distance/speed
	fmt.Println(time, "hours", time/24, "days")
}