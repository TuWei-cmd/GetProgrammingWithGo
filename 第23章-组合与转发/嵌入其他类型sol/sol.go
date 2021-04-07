package main

import "fmt"

// 摄氏度
type celsius float64

// 日期
type sol int

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

// 位置
type location struct {
	lat, long float64
}

func (l location) days(l2 location) int {
	//待办事项：复杂的距离运算
	return 5
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
	sol
	temperature
	location
}

func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

func main() {
	report := report{
		sol:         15,
		temperature: temperature{high: -1.0, low: -78.0},
		location:    location{-4.5895, 137.4417},
	}
	fmt.Println(report.sol.days(1446))
	fmt.Println(report.days(1446))
}
