package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %v°C\n", temp)
	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the Moon is %v°C\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
}
