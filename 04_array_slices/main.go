package main

import "fmt"

func main() {
	// ARRAY - have to be FIXED length, name the types
	nums := [2]int{23, 16}
	var names [2]string

	// assign values
	names[0] = "a"
	names[1] = "b"

	fmt.Println(nums)     // [23 16]
	fmt.Println(names[0]) // [a]

	// SLICES
	colors := []string{"pink", "red", "yellow", "orange"}

	fmt.Println(colors)      // [pink red yellow orange]
	fmt.Println(len(colors)) // get the length of slices
	fmt.Println(colors[1:3]) // [red yellow]
}
