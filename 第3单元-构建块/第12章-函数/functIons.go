package main

import (
	"fmt"
)


func main() {
	celsius := 18.6 
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Print(celsius, " C is ", fahrenheit, " F.")
}

//K氏度转摄氏度
func kelvinToCelsius(k float64) float64 {
	k -= 237.15
	return k
}

//摄氏度转华氏度
func celsiusToFahrenheit(c float64) float64 {
	return (c *9.0/5.0)+32.0
}

//K氏度转华氏度
func kelvinToFarenheit(k float64) float64 {
	k -= 237.15
	return (k *9.0/5.0)+32.0
}