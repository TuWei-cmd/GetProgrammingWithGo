// 定制Location结构中的字段
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	curiosity := location{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(bytes)
	fmt.Println(string(bytes))
}

// exitOnError打印所有错误和退出信息
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
