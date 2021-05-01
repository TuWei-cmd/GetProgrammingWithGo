package main

import "fmt"

func main() {
	// 接口可以为nil
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil)

	// 这是啥
	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil)

	// 检视接口变量的内部表示
	fmt.Printf("%#v\n", v)
}
