package main

import "fmt"

// 动物类型，实现名字，食物，运动
type animal struct {
	name   string
	foods  []string
	moving string
}

// 实现接口Stringer,返回名字
func (a animal) String() string {
	return fmt.Sprintf("%v", a.name)
}

// 描述食物
func (a animal) eat() string {
	return fmt.Sprintf("%v eat %v.", a, a.foods)
}

// 描述运动
func (a animal) move() string {
	return fmt.Sprintf("%v move by %v.", a, a.moving)
}

func main() {
	horse := animal{
		name:   "HORSE",
		foods:  []string{"grass"},
		moving: "running",
	}
	fmt.Println(horse)
	fmt.Println(horse.eat(), horse.move())
}
