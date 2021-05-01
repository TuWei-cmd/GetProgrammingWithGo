package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrBounds = errors.New("out of bounds")
	Errdigit  = errors.New("out of digit")
)

const rows, columns = 9, 9

// Grid是一个数独网络
type Grid [rows][columns]int8

// 下棋：检查形参
func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return ErrBounds
	}
	if !validDigit(digit) {
		return Errdigit
	}
	g[row][column] = digit
	return nil
}

// 辅助函数:检查下棋点
func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

// 辅助函数:检查点数
func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func main() {
	var g Grid
	err := g.Set(12, 3, 5)
	if err != nil {
		switch err {
		case ErrBounds, Errdigit:
			fmt.Println("Les erreurs de parametres hors limites.")
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}
}
