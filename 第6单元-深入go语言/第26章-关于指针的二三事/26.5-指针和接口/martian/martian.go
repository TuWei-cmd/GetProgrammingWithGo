// 指针和接口
package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

func main() {
	// 无论是martian还是指向martian的指针，都满足talker接口
	shout(martian{})
	shout(&martian{})

	// 如果方法使用的是指针接收者，情况会有所不同
	pew := laser(2)
	shout(&pew)
	// shout(pew)
}
