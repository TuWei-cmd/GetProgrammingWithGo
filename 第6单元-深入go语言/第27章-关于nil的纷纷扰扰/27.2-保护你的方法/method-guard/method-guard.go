package main

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {
	// 防范惊恐
	if p == nil {
		return
	}
	p.age++
}

func main() {
	var nobody *person
	fmt.Println(nobody)
	nobody.birthday()
	fmt.Println(nobody)
}
