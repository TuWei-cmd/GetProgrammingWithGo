package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v \n", label, len(slice), cap(slice))
}

func main() {
	dwarfs := make([]string, 0, 2)
	capDwarfs := cap(dwarfs)
	dump("dwarfs", dwarfs)
	for i := 0; i < 20; {
		element := "elmt" + fmt.Sprintf("%v", i)
		dwarfs = append(dwarfs, element)
		if cap(dwarfs) != capDwarfs {
			capDwarfs = cap(dwarfs)
			dump("dwarfs", dwarfs)
			i++
		}
	}

}
