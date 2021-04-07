package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

// 大声喊出想说的话
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

//通过嵌入满足接口
type starship struct {
	laser
}

func main() {
	// 大声呼喊
	shout(martian{})
	shout(laser(5))
	s := starship{laser(9)}
	shout(s)
	// 多态
	var t talker
	t = martian{}
	fmt.Println(t.talk())
	t = laser(3)
	fmt.Println(t.talk())
}
