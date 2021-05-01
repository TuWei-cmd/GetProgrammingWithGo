package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"
func main() {
	for i:=5; i>0; i--{
		year := 2018
		month:=rand.Intn(12)+1
		dayInMonth := 31
		switch month {
		case 2 :
			dayInMonth = 28
		case 4, 6, 9, 11:
			dayInMonth = 30
		}
		day := rand.Intn(dayInMonth) + 1
		fmt.Println(era, year, month, day)
	}
}