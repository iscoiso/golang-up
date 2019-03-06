package main

import (
	"fmt"
	"strconv"
)

/*
	MAIN TYPES:
		- string
		- bool
		- int
		- int int8 int16 int32 int64
		- uint uint8 uint16 uint32 uint64 uintptr
		- byte - alias for uint8
		- rune - alias for int32
		- float32 float64
		- complex64 complex128

*/

func main() {
	var name string = "Ismet Kovacevic"
	var age int = 22
	const isCool bool = true
	age = 23

	// short way
	username := "ismetkov"
	email, token := "kovacevicismet@gmail.com", "dsad12311l2pe12le12pld1d2"

	fmt.Println("My name is " + name + " I have " + strconv.Itoa(age))
	fmt.Printf("%T\n", name) // get the type
	fmt.Println(username, email, token)
}
