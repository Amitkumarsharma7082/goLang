package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	// defer LIFO order
	data := add(3, 5)
	defer fmt.Println("Data is :", data)
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}
