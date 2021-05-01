package main

import (
	"fmt"
)

type celsius float64
type kelvin float64
type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

// celsiusToKelvin converts °C to °K
func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var c celsius = 294.0
	k := c.kelvin()
	fmt.Print(k, "°K is ", c, "°C")
}
