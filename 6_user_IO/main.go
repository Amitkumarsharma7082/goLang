package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello enter name")

	reader := bufio.NewReader(os.Stdin) // its create a new reader that read from Stdin
	name, _ := reader.ReadString('\n')
	fmt.Println("My name : ", name)

}

// func main() {
// 	var n int

// 	for {
// 		fmt.Scan(&n)
// 		if n == 42 {
// 			break
// 		}
// 		fmt.Println(n)
// 	}

// }

// func main() {
// 	var name string
// 	fmt.Print("Enter the name : ")
// 	fmt.Scan(&name) // read until whitespace
// 	fmt.Printf("Hello %s!\n", name)
// }

// func main() {
// 	var firstName, lastName string
// 	var age int

// 	fmt.Print("Enter your first name, last name and age : ")
// 	fmt.Scan(&firstName, &lastName, &age)
// 	fmt.Printf("Hello %s %s, age %d\n", firstName, lastName, age)
// }
