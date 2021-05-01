package main

import "fmt"

type location struct {
	lat  float64
	long float64
}

func main() {
	var curiosity = location{lat: -4.5895, long: 137.4417}

	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var opportunity location
	opportunity.lat = -1.9462
	opportunity.long = 354.4734
	fmt.Println(curiosity, spirit, opportunity)
	fmt.Printf("%+v %+v %+v\n", curiosity, spirit, opportunity)
}
