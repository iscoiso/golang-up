package main

import "fmt"

func main() {
	// Pointers - allow you to point to the memory address/location of a value that is a variable
	a := 6
	b := &a

	fmt.Println(a, b)

	// use * to read val from address
	fmt.Println(*b)

	// change val with pointer
	a = 16
	fmt.Println(*b)
}
