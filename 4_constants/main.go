package main

import "fmt"

const isAdult = true

func main() {
	const name = "golang"
	// name = "JS" error
	fmt.Println(name)

	var age = 30
	age = 32
	fmt.Println(age)

	fmt.Println(isAdult)

	//* const grouping
	const (
		Port = 5000
		host = "locahost"
	)

	fmt.Println(Port, host)

}
