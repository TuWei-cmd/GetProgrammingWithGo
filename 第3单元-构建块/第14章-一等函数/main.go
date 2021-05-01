package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

// 测温
func measureTemperature(samples int, s sensor) {
	for i := 0; i < samples; i++ {
		k := s()
		fmt.Printf("%v° K\n", k)
		time.Sleep(time.Second)
	}
}

// 虚拟传感器
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

// 真实传感器
func realSensor() kelvin {
	return 0
}

// 传感器校准
func calibrates(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var offset kelvin = 3
	ss := calibrates(fakeSensor, offset)
	fmt.Println(ss(), ss(), ss(), ss())
}
