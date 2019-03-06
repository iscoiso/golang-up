package main

import "fmt"

func sayHi(name string) string {
	return "Hi, my name is " + name
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	msg := sayHi("Ismet")
	total := sum(2, 2)

	fmt.Println(msg)
	fmt.Println(total)
}
