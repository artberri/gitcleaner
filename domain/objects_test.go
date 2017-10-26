package domain

import (
	"fmt"
	"sort"
)

func ExampleBySizeDesc() {
	object := []GitObject{
		{"87a6asmall", 100, 500, "./"},
		{"87a6axxl", 31500, 500, "./"},
		{"87a6abig", 2100, 500, "./"},
		{"87a6amedium", 800, 500, "./"},
	}

	sort.Sort(BySizeDesc(object))
	fmt.Println(object)
	// Output:
	// [{87a6axxl 31500 500 ./} {87a6abig 2100 500 ./} {87a6amedium 800 500 ./} {87a6asmall 100 500 ./}]
}
