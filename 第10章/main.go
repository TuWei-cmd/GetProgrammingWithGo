package main

import (
	"fmt"
	"math"
	"strconv"
)

func mars_Age() {
	fmt.Println("计算作者在火星上的年龄：")
	age := 41 
	marsAge := float64(age)
	
	marsDays := 687.0 
	earthDays := 365.2425 
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("I am", int(marsAge), "years old on Mars.")
}
func ariane() {
	fmt.Println("阿丽亚娜类型转换：")
	var bh float64 = 32768
	if bh < math.MinInt16 || bh > math.MaxInt16 {
		defer fmt.Println("bh超出int16的范围！")
	}
	var h = int16(bh)
	fmt.Println(h)
}
func runeConvert() {
	fmt.Println("将类型int32转换位string:")
	var pi int32 = 960
	var alpha int32 = 940
	var omega int32 = 969
	var bang int32 = 33
	fmt.Println(string(pi),string(alpha),string(omega),string(bang) )
}
func itoa() {
	fmt.Println("将整数转换为ASCII字符：")
	countdown := 10
	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds."
	fmt.Println(str)

	str2 := fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	fmt.Println(str2)

	//字符串转换为数值
	fmt.Println("字符串转换为数值")
	countdown2, err := strconv.Atoi("-12")
	if err != nil {
		fmt.Println("哎呀，有地方出错了！")
	}
	fmt.Println(countdown2)
}
func launch() {
	fmt.Println("将布尔值转换为字符串：")

	launch := false

	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText)
	
	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch:", yesNo)
}
func tobool() {
	fmt.Println("将字符串转换为布尔值：")
	yesNo := "no"
	launch := (yesNo == "yes")
	fmt.Println("Ready for launch:", launch)
}

func main() {
	launch()
}
