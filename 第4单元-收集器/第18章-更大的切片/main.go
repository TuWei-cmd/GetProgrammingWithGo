package main

import (
	"fmt"
)

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs := make([]string, 0, 10)
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	dump("dwarfs", dwarfs)
	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna", "Orcus")
	dump("dwarfs", dwarfs)
	dump("dwarfs[4:7]", dwarfs[4:7])
}
