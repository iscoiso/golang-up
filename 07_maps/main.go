package main

import "fmt"

func main() {
	// maps are key value pairs. its like objects in javascript

	// define map
	users := map[string]string{"ipe": "ipelix@gmai.com", "alex": "alex@gas.com"}
	emails := make(map[string]string)

	emails["bob"] = "bo@gmail.com"
	emails["alex"] = "alex@gmail.com"

	// delete from map
	delete(emails, "bob")
	fmt.Println(emails) // map[bob:bo@gmail.com alex:alex@gmail.com]
	fmt.Println(users)
}
