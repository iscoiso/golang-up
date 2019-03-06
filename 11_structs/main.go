package main

import (
	"fmt"
	"strconv"
)

// define struct
type Person struct {
	name string
	age  int
}

// value reciever
func (p Person) sayHi() string {
	return "Hi, its a " + p.name + " and I'm " + strconv.Itoa(p.age)
}

// increase age - pointer reciever
func (p *Person) increase() {
	p.age++
}

func main() {
	person := Person{name: "ismetkov", age: 22}

	person.increase()
	fmt.Println(person.sayHi())
}
