package main

import "fmt"

func main() {
	var administrator *string

	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator)

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)

	// 修改指针指向的变量
	fmt.Println("// 修改指针指向的变量")
	bolden = "Charles Frank Bolden Jr."
	fmt.Println(*administrator)

	// 通过解引用administrator来间接改变bolden的值也是可以的：
	fmt.Println("// 通过解引用administrator来间接改变bolden的值也是可以的：")
	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(bolden)

	// 把administrator赋值给major将产生一个同样指向bolden的字符串指针
	fmt.Println("// 把administrator赋值给major将产生一个同样指向bolden的字符串指针")
	major := administrator
	*major = "Major General Charles Frank Bolden Jr."
	fmt.Println("administrator == major:", administrator == major)
	fmt.Println("bolden", bolden)

	// Charles Bolden的后任者Robert M. Lightfoot Jr.于2017年1月20日开始任职
	Lightfoot := "Robert M. Lightfoot Jr."
	administrator = &Lightfoot
	fmt.Println("// Charles Bolden的后任者Robert M. Lightfoot Jr.于2017年1月20日开始任职")
	fmt.Println("administrator == major:", administrator == major)

	// 把解引用major的结果赋值给另一个变量将产生字符串的一个副本。在克隆完成后，直接或者间接修改bolden将不会影响到charles的值，反之亦然：
	fmt.Println("// 把解引用major的结果赋值给另一个变量将产生字符串的一个副本。在克隆完成后，直接或者间接修改bolden将不会影响到charles的值，反之亦然：")
	charles := *major
	*major = "Charles Bolden"
	fmt.Println("charles:", charles)
	fmt.Println("bolden:", bolden)

	// 正如接下来这段代码中的charles和bolden所示，即使两个变量持有不同的内存地址，但只要它们包含的字符串相同，它们就是相等的
	fmt.Println("\n// 正如接下来这段代码中的charles和bolden所示，即使两个变量持有不同的内存地址，但只要它们包含的字符串相同，它们就是相等的")
	charles = "Charles Bolden"
	fmt.Println("charles == bolden", charles == bolden)
	fmt.Println("&charles == &bolden", &charles == &bolden)
}
