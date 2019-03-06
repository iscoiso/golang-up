package main

import "fmt"

func main() {
	ids := []int{2, 5, 12, 6, 88, 1, 3}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// add ids together
	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println(sum)
}
