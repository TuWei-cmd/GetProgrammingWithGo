// 通向惊恐的nil
package main

import "fmt"

func main() {
	// nil指针会引发惊恐
	var nowhere *int
	fmt.Println(nowhere)
	// fmt.Println(*nowhere)

	// 防范惊恐
	if nowhere != nil {
		fmt.Println(*nowhere)
	}
}
