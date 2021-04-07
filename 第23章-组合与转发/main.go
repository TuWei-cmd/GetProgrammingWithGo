package main

import (
	"fmt"
)

// 摄氏度
type celsius float64

// 位置
type location struct {
	lat, long float64
}

// 温度
type temperature struct {
	high, low celsius
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

// 位于结构中的结构
type report struct {
	sol         int
	temperature temperature
	location    location
}

// 求report中temperature的平均值
func (r report) average() celsius {
	return r.temperature.average()
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", report)

	fmt.Printf("a balmy %v° C\n", report.temperature.high)

	fmt.Printf("The average temperature is %v", report.average())
}
