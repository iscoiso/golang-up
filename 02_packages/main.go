package main

import (
	"fmt"

	"github.com/ismetkov/hello_world/02_packages/rvrs"
)

func main() {
	reversedName := rvrs.Reverse("ismetkov")

	fmt.Println(reversedName)
}
