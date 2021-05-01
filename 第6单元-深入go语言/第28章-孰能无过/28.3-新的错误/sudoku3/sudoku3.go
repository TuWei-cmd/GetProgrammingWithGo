package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrBounds = errors.New("out of bounds")
	Errdigit  = errors.New("invalid digit")
)

type SudokuError []error

// Error方法返回一个或者多个以逗号分隔的错误
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

const rows, columns = 9, 9

// Grid是一个数独网络
type Grid [rows][columns]int8

// 下棋：检查形参
func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, Errdigit)
	}
	if len(errs) > 0 {
		return errs
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
	err := g.Set(12, 3, 100)
	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
		os.Exit(1)
	}
}
