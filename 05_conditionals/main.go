package main

import "fmt"

func main() {
	x := 10
	y := 2

	// if/else, ==, <, >, <=, >=, &&, ||
	if x == y {
		fmt.Printf("%d is greater/equal %d\n", x, y)
	} else {
		fmt.Printf("%d is grater %d\n", y, x)
	}

	// else if
	color := "pink"

	if color == "pink" {
		fmt.Println("very, veryyyy cool color!")
	} else if color == "red" {
		fmt.Println("red is okay, but...")
	} else {
		fmt.Println("haaa, come on")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("red is red")
	case "pink":
		fmt.Println("most fancies color :)))")
	default:
		fmt.Println("ahh, what to say about that?")
	}
}
