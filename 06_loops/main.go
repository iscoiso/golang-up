package main

func main() {
	// long method
	i := 1

	for i <= 10 {
		println(i)
		i++
	}

	// short way
	for i := 1; i <= 10; i++ {
		println(i)
	}
}
