package main

import "fmt"
import "time"

func wraps() {
	fmt.Println("整数回绕：")
	var red uint8 = 255 
	red++
	fmt.Println(red)

	var number int8 = 127
	number++
	fmt.Println(number)
}

func bits() {
	fmt.Println("打印二进制数：")
	var green uint8 = 3
	fmt.Printf("%08b\n", green)
	green++
	fmt.Printf("%08b\n", green)
}

func bitsWrap() {
	fmt.Println("二进制在整数回绕时的状态：")
	var blue uint8 = 255
	fmt.Printf("%08b\n", blue)
	blue++
	fmt.Printf("%08b\n",blue)
}

func timeUnix() {
	fmt.Println("使用64位整数储存时间：")
	future := time.Unix(12622780800, 0)
	fmt.Println(future)
}

func main() {
	timeUnix()
}