package main

import "fmt"

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

// 获取平均温度
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

// 天气报告
type report struct {
	sol int
	temperature
	location
}

func main() {
	report := report{
		sol:         15,
		temperature: temperature{high: -1.0, low: -78.0},
		location:    location{-4.5895, 137.4417},
	}
	fmt.Printf("average %v°C\n", report.average())
	fmt.Printf("high: %v°C\n", report.high)
}
