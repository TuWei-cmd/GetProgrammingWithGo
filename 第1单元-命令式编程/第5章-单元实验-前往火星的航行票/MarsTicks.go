package main 

import (
	"fmt"
	// "math"
	"math/rand"
)

func main() {
	fmt.Printf("%-16v %-15v %-15v %v\n", "太空旅行公司", "飞行天数", "飞行类型", "价格（百万美元）")
	var (
		cops = [3]string{"SpaceX", "Space Adventures", "Virgin Galactic"}
		ways = [2]string{"单程", "往返"}
		days int // 天数
		prices int // 价格
		cop int // 公司选择
		way int // 是否往返
	)
	for i:=10; i>0; i-- {
		days = rand.Intn(22)+20
		prices = days*2 + cop*3 + rand.Intn(4) - 5
		cop = rand.Intn(3) 
		way = rand.Intn(2)

		fmt.Printf("%-25v %-18v %-20v %v\n", cops[cop], days, ways[way], prices*(way+1))
	}
}
