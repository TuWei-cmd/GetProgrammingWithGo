package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("0第6单元-深入go语言/第28章-孰能无过/28.1-处理错误/file/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
