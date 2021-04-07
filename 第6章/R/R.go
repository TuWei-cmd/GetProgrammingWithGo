package main

import (
	"fmt"
)

func main() {
	selection := 4

	switch selection {
	case 1:
		fmt.Println("格式化打印浮点数：")
		third := 1.0 / 3
		fmt.Println(third)
		fmt.Printf("%v\n", third)
		fmt.Printf("%f\n", third)
		fmt.Printf("%.3f\n", third)
		fmt.Printf("%4.2f\n", third)
		fmt.Printf("%05.2f\n", third)
	case 2:
		fmt.Println("浮点数不精确的例子：")
		piggyBank := 0.1
		piggyBank += 0.2
		fmt.Println(piggyBank)
	case 3:
		fmt.Println("先执行除法算法：")
		celsius := 21.0
		fmt.Print((celsius/5.0*9.0)+32, " F\n")
		fmt.Print((9.0/5.0*celsius)+32, " F\n")
	case 4:
		fmt.Println("先执行乘法运算：")
		celsius := 21.0
		fahrenheit := (celsius * 9.0 / 5.0) +32.0
		fmt.Print(fahrenheit, " F\n")
	}

}